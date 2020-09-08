package main

import (
	"fmt"
	//	"io/ioutil"
	"bufio"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("ss")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println(stat.Size())
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			panic(err)
		}
		line = strings.TrimSpace(line)
		if strings.Contains(line, "tag-list") {
			fmt.Println(line)
		}
	}
}
