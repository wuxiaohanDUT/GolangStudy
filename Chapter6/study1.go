package main

/**
方法声明、指针接收者方法
*/
import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Number int

func main() {
	p := Point{3.0, 4.0}
	fmt.Println("Before func call ", p)
	p.Reverse()
	fmt.Println("After func call", p)
	p.ReverseByPointer() //编译器会隐式地获取变量的地址 等同于 (&p).ReverseByPointer()
	fmt.Println("After func call by use pointer", p)
	var p1 *Point
	p1.IsNilPoint()
}

func (p Point) Distance(q Point) float64 {
	//X := 12.0 与Point处于同一命名空间
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//与包处于不同命名空间
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//可以将方法绑定到任意类型上
func (n Number) Abs() Number {
	return Number(math.Abs(float64(n)))
}

//绑定类型对象也采用值传递的方式
func (p Point) Reverse() {
	p.X, p.Y = p.Y, p.X
}
func (p *Point) ReverseByPointer() {
	p.X, p.Y = p.Y, p.X
}

//nil是一个合法的接收者
func (p *Point) IsNilPoint() {
	if p == nil {
		fmt.Println("Is a nil point")
	} else {
		fmt.Println("Is not a nil point")
	}
}
