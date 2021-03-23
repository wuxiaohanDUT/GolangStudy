package main

/**
sort的使用、自定义排序
*/
import (
	"fmt"
	"sort"
)

//自定义排序方法,定义一个结构体实现对应的接口Interface
type Person struct {
	Name string
	Age  int
}
type Persons []Person

func (p Persons) Len() int {
	return len(p)
}
func (p Persons) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}
func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func main() {
	var arr = [5]int{1, -2, 6, 7, 0}
	sort.Ints(arr[:])
	fmt.Println(arr)
	var persons Persons = []Person{
		{Name: "w", Age: 17},
		{Name: "x", Age: 15},
		{Name: "h", Age: 18},
		{Name: "n", Age: 11},
	}
	sort.Sort(persons)
	fmt.Println(persons)
}
