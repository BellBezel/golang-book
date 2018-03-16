package main

import (
	"fmt"
	"net/http"
	//"time"
	"github.com/gorilla/mux"
	//"encoding/json"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city := vars["city"]
	if city == "hobart" {
		fmt.Fprintln(w, "Hobart")
		fmt.Fprintln(w, "14c shower rain")
	}

	if city == "all" {
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

//http://localhost:8882/api/v1/weather/{city}

func APIget(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city := vars["hobart"]
	url := "localhost:8882/api/v1/weather/"
	res, _ := http.Get(url + city)

	/*type User struct {
		getCity string
	}

	type APIHandler struct{}

	func (h *APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		user := new(User)
		json.NewDecoder(r.Body).Decode(user)
		//user.CreatedAt = time.Now()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		data, _ := json.Marshal(user)
		w.Write(data)
	}*/
}

/*func UsersHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Users Page")
}*/

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/weather/{city}", HomePageHandler).Methods("GET")
	r.HandleFunc("/api/v1/weather/{city}", APIget).Methods("GET")
	//r.HandleFunc("/users", UsersHandle).Methods("GET")
	return r
}

func main() {
	http.ListenAndServe(":3000", NewRouter())
}
