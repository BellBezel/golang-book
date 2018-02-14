package main

import "fmt"

func main() {
	arr := [5]int{1,2,3,4,5}
	fmt.Println(arr)

	slice := arr[0:3]    //arr[start:end(+1)]
	fmt.Println(slice)
}