package main

import (
	"encoding/json"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

type response struct {
	Payload string `json:"payload"`
}

func register(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	if email == "" {
		jsonResponse(response{"Missing email"}, w)
		return
	}

	user, err := insertUser(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := generateToken(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse(response{token}, w)
}

var private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("user")
	id := token.(*jwt.Token).Claims.(jwt.MapClaims)["id"].(string)

	user, err := findUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse(response{user.Email}, w)
})

func jsonResponse(r response, w http.ResponseWriter) {
	jsonResponse, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
