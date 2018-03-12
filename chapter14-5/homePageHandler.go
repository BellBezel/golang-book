package home

import (
	"time"
	"net/http"
	"encoding/json"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json: "email"`
	CreatedAt time.Time `json: "created_at"`
}

type HomePageHandler struct{}
	
func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	json.NewDecoder(r.Body).Decode(user)
	user.CreatedAt = time.Now()
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	
	data, _ := json.Marshal(user)
	w.Write(data)
}

/* เข้า Postman: http://localhost:3000/
INPUT : {
	"firstName": "espresso",
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
