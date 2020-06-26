package groutine_test

import (
	"fmt"
	"testing"
)

func TestGrouting(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func TestGrouting1(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}
