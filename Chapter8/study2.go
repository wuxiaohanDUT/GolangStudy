package main

/**
无缓冲通道
*/
import "fmt"

func main() {
	flag := false
	ch1 := make(chan int) //无缓冲区通道
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
