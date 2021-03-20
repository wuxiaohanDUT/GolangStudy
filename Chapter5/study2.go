package main

/**
多返回值函数、裸返回函数
*/
import (
	"fmt"
)

func main() {
	var a, b = 3, 4
	fmt.Println(AddAndSub(a, b))
	fmt.Println(SubAndAdd(a, b))
	var c, d = AddAndSub(a, b)
	var e, _ = SubAndAdd(a, b) //可以忽略其中某个返回值
	fmt.Println(c, d, e)

}

//多返回值
func AddAndSub(a, b int) (int, int) {
	return a + b, a - b
}

//裸返回
func SubAndAdd(a, b int) (c, d int) {
	c = a - b
	d = a + b
	return
}
