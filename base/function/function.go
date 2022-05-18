package main

import "fmt"

func main() {
	var result = max(1, 10)
	fmt.Printf("max %d", result)
	fmt.Println()
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func swap(x, y string) (string, string) {
	return y, x
}
