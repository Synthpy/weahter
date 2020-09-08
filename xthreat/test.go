package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	f, _ := os.Open("test")
	s, _ := ioutil.ReadAll(f)
	defer f.Close()
	ct := string(s)
	if strings.Contains(ct, "扫描") {
		fmt.Println("true sm")
	}
	fmt.Println(s)
	fmt.Println(ct)
}
