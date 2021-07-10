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
	//a := []string{"10000000000000000000000001"}
	a := strings.Split("1603", "")
	b := strings.Split("6308", "")
	//	fmt.Println(reflect.TypeOf(a))

	al := sliceToList(a)
	bl := sliceToList(b)

	fmt.Println(al)
	fmt.Println(bl)
	//	val := []int{}
	//	vbl := []int{}

	va := al.Val
	vb := bl.Val
	fmt.Println(va, vb)
	//	val = append(val, va)
	//vbl = append(vbl, vb)

	var res []int
	jws := 0
	res = append(res, (va+vb+jws)%10)
	jws = (va + vb + jws) / 10

	fmt.Println(res)

	for {

		if al.Next != nil {
			al = al.Next
			va = al.Val
		} else {
			va = 0
		}
		if bl.Next != nil {
			bl = bl.Next
			vb = bl.Val
		} else {
			vb = 0
		}

		if (va+vb+jws) == 0 && bl.Next == nil && al.Next == nil {

		} else {
			res = append(res, (va+vb+jws)%10)
		}
		jws = (va + vb + jws) / 10
		fmt.Println(jws, "jws")
		//	val = append(val, va)
		//vbl = append(vbl, vb)
		if (va+vb+jws) == 0 && bl.Next == nil && al.Next == nil {
			break
		}

	}

	fmt.Println(va, vb, jws)
	fmt.Println("------------")
	fmt.Println(res)
	//	fmt.Println(reversInts(val))
	fmt.Println("------------")
	//	valn := sliceToNun(val)
	//vbln := sliceToNun(vbl)
	//fmt.Println(valn + vbln)
	// c := addTwoNumbers(al, bl)
	// s := ListToSlice(c)
	// fmt.Println(s)
	// aa := 123
	// fmt.Println(reverseNum(aa))

}

func reversInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reversInts(input[1:]), input[0])

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
		sum = sum + t
	}
	return sum
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	al := l1
	bl := l2
	val := []int{}
	vbl := []int{}

	va := al.Val
	vb := bl.Val
	fmt.Println(va, vb)
	val = append(val, va)
	vbl = append(vbl, vb)

	for {

		if al.Next != nil {
			al = al.Next
			va = al.Val
		} else {
			va = 0
		}
		if bl.Next != nil {
			bl = bl.Next
			vb = bl.Val
		} else {
			vb = 0
		}
		fmt.Println(va, vb)
		val = append(val, va)
		vbl = append(vbl, vb)
		if bl.Next == nil && al.Next == nil {
			break
		}

	}
	fmt.Println("------------")
	fmt.Println(val, vbl)
	fmt.Println(reversInts(val))
	fmt.Println("------------")
	valn := sliceToNun(val)
	vbln := sliceToNun(vbl)
	fmt.Println(valn + vbln)
	return numToList(valn + vbln)
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
