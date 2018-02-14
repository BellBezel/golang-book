package main

import "fmt"
import "strconv"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(FizzBuzz(i))
	}
}

func FizzBuzz(i int) string {
	if i%15==0 {
		return strconv.Itoa(i) + " FizzBuzz"
	} else if i%3==0 {
		return strconv.Itoa(i) + " Fizz"
	} else if i%5==0 {
		return strconv.Itoa(i) + " Buzz"
	} else {
		return strconv.Itoa(i)
	} 
}

/*func FizzBuzz(i int) (int,string) {
	if i%15==0 {
		return i,"FizzBuzz"
	} else if i%3==0 {
		return i,"Fizz"
	} else if i%5==0 {
		return i,"Buzz"
	} else {
		return i, "Nothing"
	} 
}*/