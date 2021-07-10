package main

import "fmt"

func main() {
	var a = [...]string{
		1: "fdf",
		5: "fff",
	}
	fmt.Println(len(a))
	fmt.Println(a[4])

}
