package main

/**
无缓冲通道、缓冲管道
*/
import (
	"fmt"
	"time"
)

func main() {
	s1()
	s2()
	//s3()
	s4()
	s5()
}
func s1() {
	flag := false
	//无缓冲区通道,需要进行同步
	ch1 := make(chan int)
	go func() {
		for i := 1; i < 11; i++ {
			//向通道发送数据
			ch1 <- i
		}
		//关闭通道
		close(ch1)
	}()
	go func() {
		for {
			//通过标志位判断通道是否关闭
			n, ok := <-ch1
			if !ok {
				flag = true
				break
			}
			fmt.Println(n)
		}
	}()
	for !flag {
	}
}
func s2() {
	//缓冲管道，不需要同步
	ch1 := make(chan int, 10)
	flag := false
	go func() {
		for i := 1; i <= 10; i++ {
			ch1 <- i
		}
		flag = true
	}()
	go func() {
		for !flag {
		}
		for {
			n, ok := <-ch1
			if !ok {
				break
			}
			fmt.Println(n)
		}
	}()
	time.Sleep(time.Second * 3)
}
func s3() {
	//尝试向一个关闭的channel中写入数据会导致宕机
	ch1 := make(chan int, 10)
	close(ch1)
	ch1 <- 10
}
func s4() {
	//可以从已关闭的channel中读取数据，但是读到的都是零值
	ch1 := make(chan int, 10)
	close(ch1)
	for i := 1; i <= 3; i++ {
		n, ok := <-ch1
		fmt.Println(n, ok)
	}
}
func s5() {
	//channel迭代
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 1; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for x := range ch1 {
			ch2 <- x * x
		}
		close(ch2)
	}()
	for x := range ch2 {
		fmt.Println(x)
	}
}
