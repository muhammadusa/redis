package hand

import (
	// "log"
	"net/http"
	"encoding/json"
)

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
}

var users []User = []User {
	User{Id: 1, Username: "Sarvar",},
	User{Id: 2, Username: "Izzat",},
}

func Users (w http.ResponseWriter, r *http.Request){
	// log.Println("Users Ctrl")

	encoder := json.NewEncoder(w)
	encoder.Encode(users)

	writeCashe := r.Context().Value("cashe").(func(interface{}, int))

	writeCashe(users, 0)
}