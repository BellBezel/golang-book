package main

import "fmt"

func main() {
	go f(0)      //Go routine will run all line without waiting the result
	var input string		// |  If comment this two lines,
	fmt.Scanln(&input)		// |  It will not print the result
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

/* OUTPUT :
0 : 0
0 : 1
0 : 2
0 : 3
0 : 4
0 : 5
0 : 6
0 : 7
0 : 8
0 : 9
*/