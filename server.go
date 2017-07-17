package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

var session *mgo.Session

func main() {
	s, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer s.Close()
	session = s

	r := mux.NewRouter()
	r.HandleFunc("/register", register).Methods("POST")
	r.Handle("/private", jwtMiddleware.Handler(private)).Methods("GET")
	http.Handle("/", r)

	fmt.Println("Server running at: 8080")
	http.ListenAndServe(":8080", nil)
}
