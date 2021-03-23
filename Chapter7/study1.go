package main

/**
接口类型、实现接口
*/
import "fmt"

type Writer interface {
	Write(p []byte) (n int, err error)
}
type Reader interface {
	Read(p []byte) (n int, err error)
}

//组合已有接口得到新接口
type WriterAndReader interface {
	Writer
	Reader
}
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	var c1 Writer
	//c1 = c 编译错误
	c1 = &c
	fmt.Println(c1.Write([]byte("world")))
	//var c2 WriterAndReader
	//c2 = &c 编译错误，必须要实现接口中定义的所有方法
	//空类型接口
	var any interface{}
	any = 123
	any = true
	any = "hello"
	fmt.Println(any)
}
