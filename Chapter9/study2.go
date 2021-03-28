package main

/**
读写互斥锁：sync.RWMutex
*/
import (
	"fmt"
	"sync"
)

var (
	Balance1 int
	Balance2 int
	Rw       sync.RWMutex
)

func GetBalance() int {
	defer Rw.RUnlock()
	Rw.RLock()
	return Balance1 + Balance2
}
func AddBalance(n int) {
	defer Rw.Lock()
	Rw.Lock()
	Balance1 += n
	Balance2 -= n
}
func SubBalance(n int) {
	defer Rw.Unlock()
	Rw.Lock()
	Balance1 -= n
	Balance2 += n
}
func main() {
	Balance1 = 1000
	Balance2 = 1000
	go func() {
		for {
			SubBalance(10)
		}
	}()
	go func() {
		for {
			AddBalance(10)
		}
	}()
	for {
		fmt.Println(getBalance())
	}
}
