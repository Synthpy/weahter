package main

import (
	"bytes"
	"flag"
	"fmt"
	"sort"
)

var temp = flag.Int("n", 1234, "help messs for flag n")
var temps = flag.String("s", "dddd", "help messs for flag d")

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	flag.Parse()
	fmt.Println(*temp)
	fmt.Println(*temps)
	var buf *bytes.Buffer
	buf = new(bytes.Buffer)
	buf.Write([]byte("ff"))

	var p *person
	sayHi(p)
	fmt.Println("abcd" < "abce")
	var s StringSlice = StringSlice{"知道m", "知道l"}
	fmt.Println(s)
	sort.Sort(s)
	fmt.Println(s)

}

type person struct{}

func sayHi(p *person) { fmt.Println("hi") }

//func (p *person) sayHi() { fmt.Println("hi") }
