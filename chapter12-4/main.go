package main

import "fmt"
import "strconv"

func main() {
	for number := 1; number <= 100; number++ {
		fmt.Println(number, FizzBuzz(number))
	}
}

func FizzBuzz(number int) string {
	fbTemplate := func(fbnumber int, str string) func(int) (string, bool) {
		return func(n int) (string, bool){
			if n%fbnumber == 0 {
				return str, true
			}
			return "", false
		}
	}

	fbArray := [...]func(n int) (string, bool){
		fbTemplate(15, "FizzBuzz"),
		fbTemplate(5, "Buzz"),
		fbTemplate(3, "Fizz"),
	}

	for i := 0; i < len(fbArray); i++ {
		if str, ok := fbArray[i](number); ok {
			return str
		}
	} 
	return strconv.Itoa(number)
}
