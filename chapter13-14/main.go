package main

import (
	"fmt"
)

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func ping(pings chan<- string, msg string) {
	pings <- msg			// ถ้าไม่ใส่ลูกศร <- มันจะเป็นได้ทั้ง receive และ sent
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings // receive
	pongs <- msg // send
}

/* OUTPUT :
passed message
*/
