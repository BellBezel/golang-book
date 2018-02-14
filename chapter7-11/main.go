package main

import "fmt"

func main() {
	// 'c' is ASCII of each character
	for i, c := range "golang" {
		fmt.Println(i, c)
		fmt.Printf("%v\n", string(c))
	}
}