package main

import "fmt"

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	const a, b, c = 1, false, "String"

	area = WIDTH * LENGTH

	fmt.Printf("面积为 %d", area)
	println()
	println(a, b, c)

	const (
		Unkown = 0
		Female = 1
		Male   = 2
	)

	println(Unkown, Female, Male)
}
