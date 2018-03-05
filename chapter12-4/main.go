package main

import "fmt"
import "strconv"

func main() {
	for number := 1; number <= 100; number++ {
		/*var addVar func(string, string) string
		addVar = func(a, b string) string {
			return a + b
		}*/
		fmt.Println(number, FizzBuzz(number))
		//fmt.Println(addVar(strconv.Itoa(i), FizzBuzz(i)))
	}
}

func FizzBuzz(number int) string {
	/*if i%15==0 {
		return " FizzBuzz"
	} else if i%3==0 {
		return " Fizz"
	} else if i%5==0 {
		return " Buzz"
	} else {
		return strconv.Itoa(i)
	} */

	

	fizzBuzzFunc := func(n int) (string, bool) {
		if n%15 == 0 {
			return "FizzBuzz", true
		}
		return "", false
	}
	buzzFunc := func(n int) (string, bool) {
		if n%5 == 0 {
			return "Buzz", true
		}
		return "", false
	}
	fizzFunc := func(n int) (string, bool) {
		if n%3 == 0 {
			return "Fizz", true
		}
		return "", false
	}

	fbArray := [...]func(n int) (string, bool){
		fizzBuzzFunc,
		buzzFunc,
		fizzFunc,
	}

	for i := 0; i < len(fbArray); i++ {
		if str, ok := fbArray[i](number); ok {
			return str
		}
	} 
	return strconv.Itoa(number)
}
