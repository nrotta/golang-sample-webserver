package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", register).Methods("POST")
	r.Handle("/private", jwtMiddleware.Handler(private)).Methods("GET")
	http.Handle("/", r)

	fmt.Println("Server running at: 8080")
	http.ListenAndServe(":8080", nil)
}
