package main

import "fmt"

func main() {
	var a string = "Test"
	fmt.Println(a)

	// 声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	// 默认值为0
	var d int
	fmt.Println(d)

	// 默认值为false
	var e bool
	fmt.Println(e)

	// 等价于
	// var f int
	// f = 1
	f := 1
	fmt.Println(f)
}
