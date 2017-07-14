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
		jsonResponse("Missing email", w)
		return
	}

	token, err := generateToken(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse(token, w)
}

var private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("user")
	jsonResponse(token.(*jwt.Token).Claims.(jwt.MapClaims)["email"].(string), w)
})

func jsonResponse(payload string, w http.ResponseWriter) {
	jsonResponse, err := json.Marshal(response{payload})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
