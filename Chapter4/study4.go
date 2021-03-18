package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	ID      int
	Name    string
	Address string
	Salary  int
}

//结构体嵌套
type Company struct {
	Name, Position string
	Employees      []Employee
}

//匿名成员,让访问更加便捷
type Point struct {
	X, Y float64
}
type Circle struct {
	Point
	Radius float64
}

//自定义序列化键名称
type Person struct {
	Age  int8   `json:"Person_Age"`
	Name string `json:"Person_Name"`
}

func main() {
	//结构体的初始化
	var emp1 = Employee{201892244, "wuxiaohan", "China", 20000}
	var emp2 = Employee{
		ID:      201892244,
		Name:    "wuxiaohan",
		Address: "China",
		Salary:  20000,
	}
	var emp3 = Employee{ID: 201892244}
	var emp4 Employee
	fmt.Println(emp1, emp2, emp3, emp4)
	//结构体可以直接进行比较
	fmt.Println(emp1 == emp2)
	var compy = Company{
		Name:      "Tencent",
		Position:  "ShenZhen",
		Employees: []Employee{emp1, emp2, emp3},
	}
	fmt.Println(compy)
	//具有匿名成员的结构体初始化
	var c = Circle{Point{3.0, 4.0}, 5.0}
	fmt.Println(c, c.X, c.Y)
	//结构体和Json类型相互转换,只有可导出的成员才会进行转换
	data, _ := json.Marshal(compy) //转化成Json
	fmt.Printf("%s \n", data)
	var cmpy Company = Company{}
	if err := json.Unmarshal(data, &cmpy); err != nil { //转化回struct
		fmt.Println(cmpy)
	}
	//自定义序列化键名称
	var person = Person{Age: 18, Name: "wuxiaohan"}
	data2, _ := json.Marshal(person)
	fmt.Printf("%s \n", data2)
	//Json字符串反序列化
	var str = `{"Person_Name": "wuxiaohan", "Person_Age": 18}`
	var person1 Person
	json.Unmarshal([]byte(str), &person1)
	fmt.Println(person1)
}
