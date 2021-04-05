package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	b := bytes.NewBuffer([]byte("hello"))
	fmt.Println(b)
	b.Write([]byte("zzz"))
	fmt.Println(b)
	var p []byte = make([]byte, 2)
	p = []byte{2, 2}
	fmt.Println(b, "bbbbbbbbbbbbbbbbbb")
	c, _ := b.Read(p)
	fmt.Println(c)
	fmt.Println(p)
	fmt.Println(b, "bbbbbbbbbbbbbbbbbb")
	os.Stdout.Write([]byte("\n"))
	os.Stdout.Write([]byte("zhidaole"))
	os.Stdout.Write([]byte("\n"))
	var bb bytes.Buffer
	bb.Write([]byte("hello "))
	fmt.Fprintf(&bb, "abcdefghijklmn")
	fmt.Printf("%T", bb)
	fmt.Println(bb, "bb")
	bn := bb.Next(2)
	fmt.Println(bb, "bb")
	bnn := bb.Next(4)
	fmt.Println(bb, "bb")
	bp := make([]byte, 2)
	bb.Read(bp)
	fmt.Println(bp, "bp")
	fmt.Println(bb, "bb")
	fmt.Println(bn, "bn")
	fmt.Println(bnn, "bnn")
	fmt.Println(bb, "bb")
	fmt.Println("--------------------------------")

	var bbb bytes.Buffer

	fmt.Println(bbb, "bb")
	bbb.Write([]byte("abcdefghizzz"))
	fmt.Println(bbb, "bb")
	pp := []byte{2, 2}
	cc, _ := bbb.Read(pp)
	fmt.Println(pp)
	cc, _ = bbb.Read(pp)
	fmt.Println(cc)
	fmt.Println(pp)
	fmt.Println(bbb, "bb")

}
