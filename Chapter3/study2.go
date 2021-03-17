package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

//常量生成器 itoa
const (
	A = iota
	B = iota << 1
	C = iota << 1
	D = iota << 1
)

func main() {
	var s1 string = "hello world"
	fmt.Println(s1, s1[3], len(s1))
	fmt.Println("hello" + "world")
	//s1[3] = '3' go中的string类型也是不可变的
	s1 += "ni hao shi jie" //生成一个新的字符串
	fmt.Println(s1)
	var s2 string = "hello"
	var s3 string = "hello"
	fmt.Println(&s2, &s3)
	//字符 rune-int32
	var c rune = rune(s1[0])
	fmt.Println(c, unsafe.Sizeof(c))
	//使用 UFT-8 进行编码
	var s4 = "你好"
	fmt.Println(s4)
	//字符转换成字节slice
	bts := []byte(s1)
	fmt.Println(bts)
	//整型转换成字符串
	var s5 = strconv.Itoa(1024)
	fmt.Println(s5)
	fmt.Println(A, B, C, D)
}
