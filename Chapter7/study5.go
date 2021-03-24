package main

/**
类型断言，有点像Java里面的instanceof
*/
import "fmt"

type Interface1 interface {
	Method1()
}
type Interface2 interface {
	Method2()
}
type Interface3 interface {
	Interface1
	Interface2
}
type S1 struct {
	Name string
}
type S2 struct {
	Name string
}
type S3 struct {
	Name string
}

func (s S1) Method1() {}
func (s S2) Method2() {}
func (s S3) Method1() {}
func (s S3) Method2() {}
func main() {
	var i1 Interface1 = S1{"name"}
	var i2 Interface2 = S2{"name"}
	var i3 Interface3 = S3{"name"}
	f1, ok1 := i1.(S1)
	f2, ok2 := i2.(S2)
	f3, ok3 := i3.(S3)
	fmt.Println(f1, f2, f3)
	fmt.Println(ok1, ok2, ok3)
}
