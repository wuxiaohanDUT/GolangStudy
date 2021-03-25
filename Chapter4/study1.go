package main

/**
数组的初始化、修改、比较
*/
import "fmt"

func main() {
	//定长数组,数组声明的几种方式
	var arr1 = [5]int{1, 2, 3, 4, 5}
	var arr2 = [...]int{1, 2, 3}
	//var arr2 = []int{1, 2, 3}
	var arr3 [3]int = [3]int{1, 2, 3}
	var arr4 = [10]int{1: 12, 3: 18, 9: 20}
	fmt.Println(arr1, arr2, arr3, arr4)
	//foreach
	for index, n := range arr1 {
		fmt.Println(index, n)
	}
	//如果数组的元素类型是可比较的，那么数组也是可比较的
	fmt.Println(arr2 == arr3)
	//注意只有类型和长度都相同的数组才能进行比较
	//fmt.Println(arr1 == arr2) 编译错误
	fmt.Println(arr4)
	changeArr(&arr4)
	fmt.Println(arr4)
	//值传递
	changearr(arr4)
	fmt.Println(arr4)
}

//函数对参数进行值传递
func changeArr(arr *[10]int) {
	arr[0] = 1024
	arr[len(arr)-1] = 1024
}
func changearr(arr [10]int) {
	arr[0] = 64
	arr[len(arr)-1] = 64
}
