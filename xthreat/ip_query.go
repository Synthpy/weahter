package main

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("ips")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	outputs := []string{}

	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("file read ok")
				break
			} else {
				fmt.Println("read file error", err)
				return
			}
		}
		ct := getResult(line)
		if strings.Contains(ct, "扫描") || strings.Contains(ct, "垃圾邮件") || strings.Contains(ct, "暴力破解") {
			outputs = append(outputs, line+"恶意")
		} else {
			outputs = append(outputs, line+"可疑")
		}

		fmt.Println(line + ct)
		//		fmt.Println("----------------")

	}

	for _, v := range outputs {
		fmt.Println(v)
	}
	//ip := "185.216.140.171"
	//ct := getResult(ip)
	//fmt.Println(ct)
	fmt.Println("----------------")

}

func getResult(ip string) string {
	url := "https://x.threatbook.cn/nodev4/ip/" + ip
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", "rememberme=3336cad4b3d1046095c8951b8b6c388ad408b3c3|caojfv@gmail.com|1599553873707")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}
	ct := doc.Find(".tag-list").First().Text()
	return ct
}
