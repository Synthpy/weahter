package array_test

import "testing"

func TestArrayTravel(t *testing.T) {
	arr3 := []int{1, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	for _, e := range arr3 {
		t.Log(e)
	}
}
