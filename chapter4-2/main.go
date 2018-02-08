package main

import "fmt"

func main() {
	fmt.Print("Enter Fahrenheit: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := (input - 32) * 5/9
	fmt.Printf("Celsius is %.2f\n",output)
}