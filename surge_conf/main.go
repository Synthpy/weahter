package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type S_vmess struct {
	V    int    `json:"v"`
	Ps   string `json:"ps"`
	Add  string `json:"add"`
	Port int    `json:"port"`
	ID   string `json:"id"`
	Aid  int    `json:"aid"`
	Net  string `json:"net"`
	Type string `json:"type"`
	Host string `json:"host"`
	Path string `json:"path"`
	TLS  string `json:"tls"`
}

func main() {
	vmess := getVmess("https://sub.bbapk.xyz/138jpGlm4ad818bc50/V2ray")
	data := parse_vmess(vmess)
	svmess := sTos(string(data))
	fmt.Println(svmess[0])
	tmpl := template.New("tmpl1")
	tmpl.Parse("{{range .svmess}} {{.ID}} iiii \n {{end}} {{.Add}}  {{.aid}}")
	var s string
	f := bytes.NewBufferString(s)
	tmpl.Execute(f, svmess[0])
	fmt.Println(f)

}

func getVmess(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	vmess, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(vmess)

}

func parse_vmess(vmess string) []byte {
	data, err := base64.StdEncoding.DecodeString(vmess)
	if err != nil {
		fmt.Println("error:", err)
	}
	return data
}

func sTos(vmess string) []S_vmess {
	vmess = strings.ReplaceAll(vmess, "vmess://", "")
	vmess = strings.TrimSuffix(vmess, "\n")
	vmesss := strings.Split(vmess, "\n")

	data := make([]S_vmess, len(vmesss))
	for i, v := range vmesss {
		d, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			fmt.Println("error:", err)
		}

		err = json.Unmarshal([]byte(d), &data[i])
		if err != nil {
			fmt.Println("error:", err)
			fmt.Println(d)
		}
	}

	return data

}
