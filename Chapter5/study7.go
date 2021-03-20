package main

/**
宕机和恢复
*/
import (
	"fmt"
)

func main() {
	fmt.Println("before panic call")
	panictest4()
	fmt.Println("after panic call")
}

//panic会中断原控制流程,一直向上返回，延迟函数会被调用，如果没有遇见recover，程序退出
func panictest1() {
	defer func() {
		fmt.Println("this is a defer func1")
	}()
	panictest2()
	fmt.Println("after panic call1")
}
func panictest2() {
	defer func() {
		fmt.Println("this is a defer func2")
	}()
	panic("this is a panic")
	fmt.Println("after panic call2")
}
func panictest3() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("get the error")
		} else {
			panic("this is a panic")
		}
	}()
	panictest1()
	fmt.Println("after panic call3")
}
func panictest4() {
	panictest3()
	fmt.Println("after panic call4")
}
