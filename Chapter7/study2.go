package main

import "fmt"

/**
多态的实现
*/
type DoSomething interface {
	SayHello()
}
type Student struct{}
type Teacher struct{}
type Professor struct{}

func (s *Student) SayHello() {
	fmt.Println("Hello I'm a student")
}
func (s *Teacher) SayHello() {
	fmt.Println("Hello I'm a teacher")
}
func (s *Professor) SayHello() {
	fmt.Println("Hello I'm a professor")
}
func main() {
	arr := [3]DoSomething{}
	arr[0] = &Student{}
	arr[1] = &Teacher{}
	arr[2] = &Professor{}
	for i := 0; i < 3; i++ {
		arr[i].SayHello()
	}
}
