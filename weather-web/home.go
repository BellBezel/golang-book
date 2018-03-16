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
	vars := mux.Vars(r)
	name := vars["name"]
	if name == "hobart" {
		fmt.Fprintln(w, "Hobart")
		fmt.Fprintln(w, "14c shower rain")
	}

	if name == "all" {
		fmt.Fprintln(w, "Hobart")
		fmt.Fprintln(w, "14c shower rain\n")

		fmt.Fprintln(w, "New York")
		fmt.Fprintln(w, "0c broken clouds\n")

		fmt.Fprintln(w, "Kupang")
		fmt.Fprintln(w, "20c clear sky\n")

		fmt.Fprintln(w, "Nairobi")
		fmt.Fprintln(w, "16c moderate rain\n")

		fmt.Fprintln(w, "Bangkok")
		fmt.Fprintln(w, "33c few clouds\n")
	}

}

func UsersHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Users Page")
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/weather/{name}", HomePageHandler).Methods("GET")
	r.HandleFunc("/users", UsersHandle).Methods("GET")
	return r
}

func main() {
	http.ListenAndServe(":3000", NewRouter())
}
