package onec_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
	"unsafe"
)

type Singleton struct {
}

var sigleInstance *Singleton

var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create obj")
		sigleInstance = new(Singleton)
	})
	return sigleInstance
}

func TestGetSingletoneObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {

			obj := GetSingletonObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func runTask(i int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("numberof task%d", i)
}

func TestFirstResponse(t *testing.T) {
	t.Log(FirstResponse())
}
