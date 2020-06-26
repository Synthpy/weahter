package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

var cur *ListNode

func main() {
	// list := new(ListNode)
	// a := []int{2, 3, 4}
	// list = sliceToList(a)
	// s := ListToSlice(list)
	// fmt.Println(s)
	// fmt.Println(sliceToNun(s))
	a := []string{"10000000000000000000000001"}
	b := []string{"564"}
	al := sliceToList(a)
	bl := sliceToList(b)
	c := addTwoNumbers(al, bl)
	s := ListToSlice(c)
	fmt.Println(s)
	aa := 123
	fmt.Println(reverseNum(aa))

}

func sliceToList(a []string) *ListNode {
	l := new(ListNode)
	cur := l

	for _, v := range a {
		vv, _ := strconv.Atoi(v)
		cur.Next = &ListNode{vv, nil}
		cur = cur.Next
	}
	return l.Next
}

func ListToSlice(l *ListNode) []int {
	var s []int
	for {
		s = append(s, l.Val)
		if l.Next == nil {
			break
		} else {
			l = l.Next
		}
	}
	return s

}

func sliceToNun(a []int) int {
	var sum int
	for i := 0; i < len(a); i++ {

		t := int(math.Pow(float64(10), float64(i))) * a[i]
		fmt.Println(t)
		sum = sum + t
	}
	return sum
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := sliceToNun(ListToSlice(l1)) + sliceToNun(ListToSlice(l2))
	s := reverseNum(sum)
	l := sliceToList(s)
	return l
}

func numToList(a int) *ListNode {
	l := new(ListNode)
	cur := l
	s := strconv.Itoa(a)
	sls := strings.Split(s, "")
	for _, v := range sls {
		vv, _ := strconv.Atoi(v)
		cur.Next = &ListNode{vv, nil}
		cur = cur.Next
	}
	return l.Next
}

func reverseNum(a int) []string {
	s := strings.Split(strconv.Itoa(a), "")
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s

}
