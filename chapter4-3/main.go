package main

import "fmt"

func main() {
	fmt.Print("Enter feet: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input*0.3048
	fmt.Printf("Meters is %.4f\n",output)
}