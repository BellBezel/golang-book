package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello, World!")
		})
	http.ListenAndServe(":3000", nil)
}

/* เข้า Web Browser: http://localhost:3000/
OUTPUT : Hello, World!
*/
