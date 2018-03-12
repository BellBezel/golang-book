package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type HomePageHandler struct{}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	json.NewDecoder(r.Body).Decode(user)
	user.CreateAt = time.Now()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	data, _ := json.Marshal(user)
	w.Write(data)
}

func main() {
	http.Handle("/", &HomePageHandler{})
	http.ListenAndServe(":3000", nil)
}

type User struct {
	FirstName string
	LastName  string
	Email     string
	CreateAt  time.Time
}

/* เข้า Postman: http://localhost:3000/
INPUT : {
	"FirstName": "espresso", 
	"lastname": "longshot", 
	"email": "espresso@speedy.coffee"
}
OUTPUT : {
    "FirstName": "espresso",
    "LastName": "longshot",
    "Email": "espresso@speedy.coffee",
    "CreateAt": "2018-03-12T10:47:50.2475768+07:00"
}
*/
