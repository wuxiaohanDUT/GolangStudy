package main

/**
单向通道类型
*/
import "fmt"

func counter(out chan<- int) {
	for x := 0; x <= 10; x++ {
		out <- x
	}
	close(out)
}
func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
}
func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
