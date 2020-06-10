package constant_test

import (
	"testing"
)

const (
	Monday = iota + 1
	Tuesday
	Wensday
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstant(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstant1(t *testing.T) {
	a := 7 //0111
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
	t.Log(Readable, a&Writeable, Executable)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	a := []int{1, 2, 3, 4}
	t.Log(len(a))
}
