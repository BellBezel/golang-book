package main

import "fmt"

func main() {
	// Long Declaration
	var x string = "Hello, World"
	fmt.Println(x)

	var y string
	y = "Hello, World"
	fmt.Println(y)

	// Short Declaration
	// Type Inference
	z := "Hello, World"
	fmt.Println(z)
	fmt.Printf("Type: %T\n",z)

	n := 555
	fmt.Println(n)
	fmt.Printf("Type: %T\n",n)

	// print set of value
	var (
		a = 5
		b = 10
		c = 15
	)
	//const b string = "Hello"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// assign and swap value
	v1,v2 := "first", "second"
	v1, v2 = v2, v1
	fmt.Println(v1)
	fmt.Println(v2)
}


