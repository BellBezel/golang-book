package main

import "fmt"
import "strconv"

func main() {
	for num := 1; num <= 100; num++ {
		fmt.Println(num, FizzBuzz(num))
	}
}

func FizzBuzz(num int) string {
	ln := [3]int{15,3,5}
	str := [3]string{"FizzBuzz","Fizz","Buzz"}
	for i := 0; i < len(ln); i++ {
		if num%ln[i]==0 {
			return str[i]
		}
	}
	return strconv.Itoa(num)
}
