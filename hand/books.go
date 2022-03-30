package hand

import (
	"net/http"
	"encoding/json"
)

type Book struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

var books []Book = []Book {
	Book{Id: 1, Name: "GO",},
	Book{Id: 2, Name: "C",},
}

func Books (w http.ResponseWriter, r *http.Request){

	encoder := json.NewEncoder(w)
	encoder.Encode(books)

	writeCashe := r.Context().Value("cashe").(func(interface{}, int))

	writeCashe(books, 0)
}