package main

/**
函数变量、匿名函数
*/
import (
	"fmt"
)

func main() {
	arr := []int{5, 6, 7, 1, -2, 3}
	sort(arr, compare1)
	fmt.Println(arr)
	sort(arr, compare2)
	fmt.Println(arr)
	//匿名函数
	sort(arr, func(a, b int) bool {
		return a-b > 0
	})
	fmt.Println(arr)
}
func compare1(a, b int) bool {
	return a-b > 0
}
func compare2(a, b int) bool {
	return b-a > 0
}

//函数参数化
func sort(arr []int, compare func(a, b int) bool) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if compare(arr[j], arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
