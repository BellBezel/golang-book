package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	//c <- 3 // overfilled the buffer
	fmt.Println(<-c)
	fmt.Println(<-c)
}

/* OUTPUT :
1
2

แต่ถ้าเปิด comment จะเกิด Deadlock เพราะเกิน buffer ที่สร้างไว้
*/
