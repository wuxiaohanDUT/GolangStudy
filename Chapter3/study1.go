package main

//go是一门强类型语言
import (
	"fmt"
)

func main() {
	//var i8 int8 = 255 + 1 溢出，编译错误
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64 = 2056
	var f64 float64 = 1024.64
	//强制类型转换
	var i = int(f64)
	var f = float64(i32)
	var s = string(i64)
	//var i = i8 + i16 编译错误，二元运算符的操作数类型必须相同
	//fmt.Println(i8 * i16) 编译错误，go中没有隐式类型转换
	fmt.Println(i8, i16, i32, i64, i, f, s)
	var p1 *int64 = &i64
	//p1 = &i32 编译错误，指针也有类型
	fmt.Println(p1)
}
