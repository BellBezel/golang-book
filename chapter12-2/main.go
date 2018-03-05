package main

import "fmt"

func main() {
	func(a, b int) int {
		fmt.Println(a+b)
		return a + b
	}(2, 3)
}