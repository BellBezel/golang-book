package main

import (
	"net/http"
	"fmt"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func UsersHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Users Page")
}


func main() {
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/users", UsersHandle)
	http.ListenAndServe(":3000", nil) 
}