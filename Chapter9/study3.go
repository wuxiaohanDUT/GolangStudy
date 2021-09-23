package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := "100"
	i1, _ := strconv.Atoi(s1)
	fmt.Println(i1)
	s2 := strconv.Itoa(i1)
	fmt.Println(s2)
	flag, _ := strconv.ParseBool("true")
	fmt.Println(flag)
	i2, _ := strconv.ParseInt("12345", 10, 64)
	fmt.Println(i2)
	s3 := strconv.FormatInt(int64(i1), 10)
	fmt.Println(s3)
}
