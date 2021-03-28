package main

/**
互斥锁：sync.Mutex
*/
import (
	"fmt"
	"sync"
)

var (
	balance1 int
	balance2 int
	mu       sync.Mutex
)

func getBalance() int {
	defer mu.Unlock()
	mu.Lock()
	return balance1 + balance2
}
func addBalance(n int) {
	defer mu.Unlock()
	mu.Lock()
	balance1 += n
	balance2 -= n
}
func subBalance(n int) {
	defer mu.Unlock()
	mu.Lock()
	balance1 -= n
	balance2 += n
}
func main() {
	balance1 = 1000
	balance2 = 1000
	go func() {
		for {
			subBalance(10)
		}
	}()
	go func() {
		for {
			addBalance(10)
		}
	}()
	for {
		fmt.Println(getBalance())
	}
}
