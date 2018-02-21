package main

import "fmt"

func main() {
	amount := 5
	double(&amount)			//Pointer
	fmt.Printf("original %v\n", amount)
}

func double(number *int) {
	*number *= 2			// '*' = call pointer
	fmt.Println(*number)
}
