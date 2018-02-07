package main

import "fmt"

func main() {
	fmt.Println("=====Zero Value=====")
	third := 1.0 / 3.0
	fmt.Printf("third = %v\n", third)
	fmt.Printf("third + third + third = %v\n", third+third+third)
}