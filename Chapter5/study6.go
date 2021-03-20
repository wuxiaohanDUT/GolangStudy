package main

/**
延迟函数调用
*/
import (
	"fmt"
	"time"
)

func main() {
	f()
	n, s := sum([]int{1, 2, 3})
	fmt.Println(n, s)
	f2()
}

//函数的延迟调用
func f() {
	defer func() {
		fmt.Println("the func end ", time.Now())
	}()
	fmt.Println("the func begin ", time.Now())
	time.Sleep(1 * time.Second)
}

//延迟调用的函数可以访问外围函数的结果，可以对其进行更改
func sum(arr []int) (total int, str string) {
	defer func() {
		total = -total
		str = fmt.Sprint("The result is ", total)
	}()
	for _, n := range arr {
		total += n
	}
	return
}

//延迟函数调用的顺序按照定义顺序的倒序进行
func f2() {
	defer func() {
		fmt.Println("f1")
	}()
	defer func() {
		fmt.Println("f2")
	}()
	fmt.Println("f")
	defer func() {
		fmt.Println("f3")
	}()
}
