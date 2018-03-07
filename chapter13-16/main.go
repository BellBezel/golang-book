package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quit <- 0
	}()
	fibonacci(ch, quit)
}

func fibonacci(ch, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// ถ้าใช้ select จะใช้ในกรณีที่ต้องการดักเรื่อง timeout ได้ (คำสั่ง close จะดักไม่ได้)

/* OUTPUT :
0
1
1
2
3
5
8
13
21
34
quit
*/
