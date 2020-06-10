package func_test

import (
	"math/rand"
	"testing"
)

func returnMultiCalues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, _ := returnMultiCalues()
	t.Log(a)
}
