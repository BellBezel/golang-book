package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	//c <- 1
	//c <- 2
	go func() { c <- 1 }()
	//go func() { c <- 2}()
	fmt.Println(<-c)
	//fmt.Println(<-c)
}

// ไม่สามารถส่งค่าให้ chan ภายใน main ได้ เพราะถ้า main ทำหน้าที่เก็บข้อมูลแล้ว ก็จะไปทำงานอื่นไม่ได้ จะต้องมี go routine มาช่วย

/* OUTPUT :
1

ถ้าเปิด comment จะเกิด Deadlock
*/
