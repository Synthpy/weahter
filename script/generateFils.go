package main

import (
	"fmt"
	"io/ioutil"
	//	"os"
)

func main() {
	dir, _ := ioutil.TempDir("/Volumes/192.168.2.143", "pre")
	fmt.Println(dir)
	for i := 0; i < 50; i++ {
		tmpfile, _ := ioutil.TempFile(dir, "example")
		s := fmt.Sprintf("%d", i)
		m := []byte(s)
		err := ioutil.WriteFile(tmpfile.Name(), m, 0644)
		if err != nil {
			panic(err)
		}
		fmt.Println(tmpfile)

	}
	files, _ := ioutil.ReadDir(".")
	fmt.Println(files)
	//	os.RemoveAll(dir)

}
