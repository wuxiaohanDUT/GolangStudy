package main

import (
	"fmt"
)

func main() {
	var n int = 1
	fmt.Println(n)
	incre(&n)
	fmt.Println(n)
	var p = test()
	fmt.Println(p)
	//new创建一个变量并初始化为零值，返回其地址
	p2 := new(int)
	fmt.Println(p2, *p2)
}

func incre(p *int) {
	*p++
}

func test() *int {
	x := 1
	return &x //逃逸,变量在堆空间中分配
}
