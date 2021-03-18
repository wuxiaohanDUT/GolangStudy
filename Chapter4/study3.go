package main

import "fmt"

func main() {
	var map1 = map[string]int{} //创建一个空map
	var map2 = map[string]int{  //创建一个非空map
		"hello":  16,
		"world":  8,
		"delete": 1,
	}
	map3 := make(map[string]int) //内置函数make创建一个map
	fmt.Println(map1, map2, map3)
	delete(map2, "delete")
	fmt.Println(map2)
	//在map中查询元素
	s1, ok1 := map2["hello"]
	s2, ok2 := map2["exits"]
	fmt.Println(s1, ok1, s2, ok2)
	//向不存在的键插入元素会导致宕机
	map2["exits"] = 1024
	//map无法直接进行比较
	//fmt.Println(map1 == map2)
	//遍历map，顺序不固定
	for i, j := range map2 {
		fmt.Println(i, j)
	}
}
