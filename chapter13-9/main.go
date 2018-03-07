package main

import (
	"fmt"
)

func main() {
	array := []int{7, 2, 8, -9, 4, 0}

	ch := make(chan int) //'make' used to create ref

	go sum(array[:len(array)/2], ch) //go routine รับค่าจาก array ตัวแรกถึงตัวกลาง แล้วส่งเข้า chan
	go sum(array[len(array)/2:], ch) //go routine รับค่าจาก array ตัวกลางถึงตัวหลัง แล้วส่งเข้า chan
	x := <-ch                        //receive value from chan (ใครมาก่อนได้ก่อน)
	y := <-ch
	fmt.Println(x, y, x+y)
}

func sum(array []int, ch chan int) {
	sum := 0
	for _, value := range array {
		sum += value
	}
	ch <- sum
}
// channal จะมีลักษณะการทำงานเหมือน Queue ไม่ได้เก็บค่าแบบถาวร
/* OUTPUT :
Round 1: -5 17 12
Round 5: 17 -5 12

*/
