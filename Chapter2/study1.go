package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	//变量的声明
	var n1 int = 10 + 20
	var s1 string = "wu"
	var r1 rune = 'c'
	var f1 float64 = 3.14
	//可以省略类型，根据表达式结果类型进行推断
	var n2 = 40
	fmt.Println(n2)
	fmt.Println(n1, s1, r1, f1)
	//短变量,变量的作用域可以被覆盖
	// := 表示声明， = 表示赋值
	temp := 1
	if true {
		var n1 int = 10
		temp := 2
		fmt.Println(temp, n1)
	}
	fmt.Println(temp)
	//同时被多个变量赋值，右边的表达式将被推演
	var n3, s2 = 123, "123"
	var n4, n5, n6 = 4, 5, 6
	n4, n5, n6 = n6+n4, n4, n5
	fmt.Println(n3, s2)
}
