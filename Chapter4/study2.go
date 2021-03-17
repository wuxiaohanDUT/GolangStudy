package main

import "fmt"

func main() {
	var arr [100]int
	for i := 0; i < 100; i++ {
		arr[i] = i
	}
	arr1 := arr[10:20]
	arr2 := arr[89:99]
	fmt.Println(arr1, arr2)
	fmt.Println(len(arr1), cap(arr1))
	//fmt.Println(arr1[999]) 超出slice cap访问，导致宕机
	//fmt.Println(arr1 == arr2) slice无法直接进行比较
	//slice数组相当于是一个对原数组的指针引用，对slice的修改最终会映射到原数组
	fmt.Println(arr)
	reverse(arr1)
	fmt.Println(arr)
	var arr4 = [5]int{0, 1, 2, 3, 4}
	arr5 := arr4[0:2]
	fmt.Println(arr5, len(arr5), cap(arr5))
	//使用 append，不产生扩容，新的slice和原slice引用相同的底层数组
	arr5 = append(arr5, 64)
	arr5[0] = 16
	fmt.Println(arr4)
	//使用 append 函数向slice中添加元素，如果容量不总会进行扩容,产生一个新的slice，新slice和原slice引用不同的底层数组
	arr5 = append(arr5, 64, 64, 64, 64)
	fmt.Println(arr5, len(arr5), cap(arr5))
	arr5[0] = 9
	fmt.Println(arr4)
	//使用make创建一个新的slice,新slice引用的不是相同的底层数组
	var arr8 = [5]int{0, 1, 2, 3, 4}
	arr7 := arr8[0:3]
	arr6 := make([]int, len(arr7), 6)
	copy(arr6, arr7)
	arr6[0] = 16
	fmt.Println(arr8)
	arr7[0] = 16
	fmt.Println(arr8)
}
func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i <= j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
