package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	/*vars := mux.Vars(r)
	name := vars["name"]

	if vars["name"] == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)*/
	fmt.Fprintln(w, "Hobart")
	fmt.Fprintln(w, "14c shower rain")
}

func UsersHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Users Page")
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/weather/hobart", HomePageHandler).Methods("GET")
	r.HandleFunc("/users", UsersHandle).Methods("GET")
	return r
}

func main() {
	http.ListenAndServe(":3000", NewRouter())
}
