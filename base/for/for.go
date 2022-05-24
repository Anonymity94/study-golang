package main

import "fmt"

func main() {
	print9x()
	printString()
	printString2()
}

func print9x() {
	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%dx%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}
}

func printString() {
	var str string = "hello world"
	for i := 0; i < len(str); i++ {
		fmt.Println()
		fmt.Printf("%d, %s", str[i], string(str[i]))
	}
}

func printString2() {
	var str string = "hello world"
	for i, v := range str {
		fmt.Printf("str[%d], %c\n", i, v)
	}
}
