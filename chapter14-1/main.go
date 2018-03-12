package main

import (
	"fmt"
	"net/http"
)

type HomePageHandler struct{}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Home!")
}

func main() {
	http.HandleFunc(
		"/", //บอกจุดเริ่มต้น ว่าจะเริ่มดูตรงนี้
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello, World!")
		})
	barHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Bar!")
	}
	http.HandleFunc("/bar", barHandler) //ถ้าใส่ /bar ที่ url จะ print "Hello Bar!"
	http.Handle("/home", &HomePageHandler{})
	http.ListenAndServe(":3000", nil) //ตรง nil เป็นจุดที่สามารถใส่สิ่งที่ต้องการค้นหาเพิ่มเติมได้ แต่ถ้าใส่ nil หมายถึงไม่ต้องการค้นหา
}

/* เข้า Web Browser: http://localhost:3000/
OUTPUT : Hello, World!

เข้า Web Browser: http://localhost:3000/bar
OUTPUT : Hello Bar!

เข้า Web Browser: http://localhost:3000/home
OUTPUT : Hello Home!
*/
