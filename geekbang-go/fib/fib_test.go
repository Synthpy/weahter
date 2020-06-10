package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {

	a := 1
	b := 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	temp := a
	a = b
	b = temp
	t.Log(a, b)

}
