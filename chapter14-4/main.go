package home

import (
	"fmt"
	"net/http"
)

func HomePageHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello, World!")
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
