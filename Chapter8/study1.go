package main

/**
goroutine的简单使用
*/
import "fmt"

func main() {
	flag := false
	go func() {
		for {
			if !flag {
				fmt.Println(1)
				flag = true
			}
		}
	}()
	go func() {
		for {
			if flag {
				fmt.Println(2)
				flag = false
			}
		}
	}()
	for {
	}
}
