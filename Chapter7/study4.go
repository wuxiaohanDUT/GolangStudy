package main

/**
error接口
*/
import (
	"errors"
	"fmt"
)

type Customerror struct {
	infoa string
	infob string
	Err   error
}

func (cerr Customerror) Error() error {
	errorinfo := fmt.Sprint("the error is %s", cerr.Err)
	return errors.New(errorinfo)
}
func main() {
	//方法一：直接采用errors包的New方法生成一个新的错误
	var err1 error = errors.New("this is a new error")
	fmt.Println(err1)
	//方法二：采用fmt.Errorf生成一个错误
	var err2 error = fmt.Errorf("this is a new error")
	fmt.Println(err2)
	//方法三：采用自定义方法实现
	var cus = Customerror{"infoa", "infob", errors.New("this is a new error")}
	var err3 error = cus.Error()
	fmt.Println(err3)
}
