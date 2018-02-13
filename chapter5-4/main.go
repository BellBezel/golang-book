package main

import "fmt"
import "math/rand"
import "time"

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	random := rand.Intn(100)
	fmt.Println("Random number is ",random)
	for i := 0; i <= 5; i++ {
		fmt.Print("Enter number : ")
		var input int
		fmt.Scanf("%d\n", &input)
		if i<5 {
			if input==random {
				fmt.Println("เจอแล้ว")
				return
			} else if input>random {
				fmt.Println("มากไป")
			} else {
				fmt.Println("น้อยไป")
			}
		} else {
			fmt.Println("เกินพอ")
			return
		}
	}
}