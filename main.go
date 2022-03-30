package main

import (
	"log"
	"net/http"

	m "app/middle"
	h "app/hand"
)


func main() {
	http.Handle("/users", m.RedisCashe(h.Users))
	http.Handle("/books", m.RedisCashe(h.Books))

	log.Println("Server ready at : 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}