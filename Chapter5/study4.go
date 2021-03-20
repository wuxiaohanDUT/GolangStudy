package main

/**
迭代变量陷阱
*/
import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	//发现迭代过程当中，变量的地址都是相同的
	for i, n := range arr {
		fmt.Println(&i, &n)
	}
	fmt.Println()
	//可以这么做,声明一个局部变量
	for i, n := range arr {
		index, num := i, n
		fmt.Println(&index, &num)
	}
}
