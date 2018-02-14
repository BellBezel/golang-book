package main

import "fmt"

func main() {
	mymap := make(map[int]int)
	mymap[1] = 1
	mymap[2] = 2

	fmt.Println(mymap[3])
	if mymap[3] != 0 {
		fmt.Println(mymap[3])
	}

	// used 'ok' to check zero value
	if value, ok := mymap[3]; ok {
		fmt.Println(value)
	}

	/* value, ok := mymap[3]
	if ok {
		fmt.Println(value)
	} */
}