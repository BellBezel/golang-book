package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc(
		"/", //บอกจุดเริ่มต้น ว่าจะเริ่มดูตรงนี้
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello, World!")
		})
	http.ListenAndServe(":3000", nil) //ตรง nil เป็นจุดที่สามารถใส่สิ่งที่ต้องการค้นหาเพิ่มเติมได้ แต่ถ้าใส่ nil หมายถึงไม่ต้องการค้นหา
}

/* เข้า Web Browser: http://localhost:3000/
OUTPUT : Hello, World!
*/
