package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)
	go func() { c <- 1 }()
	go func() { c <- 2 }()
	go func() { c <- 3 }()
	fmt.Println(<-c)
	fmt.Println(<-c)
}
// ถ้าไม่มี buffer จะทำให้เกิด data leak เพราะข้อมูลบางตัวจะไม่ถูกดึงออกมาใช้ ก็จะกินที่ไปเรื่อยๆ 
// ถ้ามี buffer ตัว go routine จะเป็นตัวจัดการให้
/* OUTPUT :
3
1
*/
