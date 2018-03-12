package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)

	for i := range ch {
		fmt.Println(i)
	}
}
11
func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)     //ถ้า comment ตัว close ไป คำสั่ง range ใน main จะไม่รู้ว่า channel นี้จบแล้ว และจะรอไปเรื่อยๆ ทำให้เกิด Deadlock
}

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
*/
