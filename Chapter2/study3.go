package main

import (
	"fmt"
)

//常量的声明
type Celsius float64
type Fahrenheit float64

//声明常量，初始化的顺序按照依赖顺序进行
const (
	A = B
	B = 1024 * C
	C = 64
)

func main() {
	var cc Celsius = 37.0
	var ff Fahrenheit = 100.0
	//var cc2 Celsius = ff 虽然底层类型相同，但是无法相互赋值，需要进行强制类型转换
	var cc2 Celsius = Celsius(ff)
	//A = 1 常量无法被修改
	fmt.Println(cc, ff, cc2)
}
