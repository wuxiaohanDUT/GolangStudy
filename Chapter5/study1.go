package main

/**
函数定义、传值方法
*/
import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var a, b = 5, 4
	fmt.Println(add(a, b), sub(a, b))
	incre(a)
	fmt.Println("after incre", a)
	realincre(&a)
	fmt.Println("after realincre", a)
	var person = Person{"wuxiaohan", 16}
	changePerson(person)
	fmt.Println("after change", person)
	var arr = []int{1, 2, 3, 4, 5}
	var slice = arr[:]
	changeSlice(slice)
	fmt.Println(slice, arr)
}

//函数定义的几种方法
func add(a, b int) int {
	return a + b
}
func sub(a int, b int) (z int) {
	z = a - b
	return
}

//函数采用的是值传递
func incre(a int) {
	a++
}
func realincre(a *int) {
	*a++
}

//结构体也采取值传递
func changePerson(person Person) {
	person.Name += "changed"
	person.Age = 18
}

//可以更改slice底层对应的数组，说明slice依靠指针实现的
func changeSlice(slice []int) {
	for i, _ := range slice {
		slice[i] = 1024
	}
}
