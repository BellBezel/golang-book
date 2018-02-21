package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}
	double(&array)
	fmt.Printf("original addr %p\n", &array)
	fmt.Printf("original %v\n", array)
}

func double(nums *[3]int) {
	fmt.Printf("double addr %p\n", nums)
	fmt.Printf("double value %v\n", *nums)
	for i := range nums {
		nums[i] *= 2
	}
	fmt.Println(*nums)
}
