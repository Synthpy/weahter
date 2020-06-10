package condition_test

import "testing"

func TestIfMultiSec(t *testing.T) {
	if _, err := somefunc(); err == nil {
		t.Log("1==1")
	} else {
		t.Log("1==1")
	}
}

func somefunc() (v int, e error) {
	return 1, nil
}

func TestSwitchMulticase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}
