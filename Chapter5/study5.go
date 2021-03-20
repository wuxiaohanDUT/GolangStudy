package main

/**
变长函数
*/
import (
	"fmt"
)

func main() {
	fmt.Printf("%T \n", sum1)
	fmt.Println(sum1(1, 3, 5, 7))
	fmt.Printf("%T \n", sum2)
	fmt.Println(sum2([]int{1, 3, 5, 7}))
}
func sum1(vals ...int) int {
	total := 0
	for n := range vals {
		total += n
	}
	return total
}
func sum2(vals []int) int {
	total := 0
	for _, n := range vals {
		total += n
	}
	return total
}
