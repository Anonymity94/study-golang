package main

import "fmt"

func main() {
	print9x()
}

func print9x() {
	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%dx%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}
}
