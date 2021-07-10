package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Need struct {
	Name string `json:"name"`
	Num  string `json:"num"`
}

type Item struct {
	Name  string `json:"name"`
	Build string `json:"build"`
	Needs []Need `json:"needs"`
}

func main() {
	r, err := ioutil.ReadFile("items.json")
	if err != nil {
		fmt.Println(err)
	}
	var items []*Item

	itemsb := bytes.Split(r, []byte("\n"))
	//fmt.Println(itemsb)
	for _, item := range itemsb {
		stt := new(Item)
		err = json.Unmarshal(item, &stt)
		if err != nil {
			fmt.Println(err)
		}
		items = append(items, stt)

	}
	for _, item := range items {
		fmt.Println(*item)
	}
	//fmt.Println(items)
	var cMap map[string]int
	cMap = make(map[string]int)
	cp := printItem(items[42], items, cMap)
	//fmt.Println(cp)
	for k, v := range cp {
		fmt.Println(k, v)
	}

}

func printItem(item *Item, items []*Item, cMap map[string]int) map[string]int {
	//fmt.Println(item.Name)
	//fmt.Println(item.Needs)
	for _, need := range item.Needs {
		num, err := strconv.Atoi(need.Num)
		if err != nil {
			fmt.Println(err)
		}
		cMap[need.Name] = cMap[need.Name] + num
		for _, v := range items {
			if need.Name == v.Name {

				for i := 0; i < num; i++ {
					//fmt.Println(num)
					printItem(v, items, cMap)
				}

			}
		}
	}
	return cMap

}
