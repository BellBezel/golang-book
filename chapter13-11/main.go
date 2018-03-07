package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3) //สร้าง chan ที่มี buffer
	go func() { ch <- 1 }()
	ch <- 2
	fmt.Println("cap:", cap(ch)) //capacity ของ chan
	fmt.Println("len:", len(ch)) //จำนวนข้อมูลที่เก็บใน chan
}

//สามารถส่งให้ chan ได้ภายใน main ได้ เพราะมีการสร้าง buffer มาช่วยเก็บ

/* OUTPUT :
cap: 3
len: 2
*/
