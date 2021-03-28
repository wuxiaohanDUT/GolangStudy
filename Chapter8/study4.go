package main

/**
select与channel配合使用
它用于等待一个或者多个channel的输出
可以用来处理异步IO操作
当多个case语句都可以被执行时，伪随机的选择一个执行
*/
import (
	"fmt"
	"time"
)

var cha chan int = make(chan int, 1)

//向channel中写入数据
func write() {
	time.Sleep(3 * time.Second)
	cha <- 1
}

//等待channel的输出，设置最久等待时间
func read() {
	select {
	case n := <-cha:
		fmt.Println(n)
		return
	case <-time.After(5 * time.Second):
		fmt.Println("read time out")
		return
	default: //所有case都被阻塞时，执行default中的语句，否则select阻塞等待
		fmt.Println()
	}
}
func main() {
	go write()
	read()
}
