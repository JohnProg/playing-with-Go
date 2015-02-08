package middleware

import (
	"../models"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError)

func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response, err := fn(w, r)

	// check for errors
	if err != nil {
		log.Printf("ERROR: %v\n", err.Error)
		http.Error(w, fmt.Sprintf(`{"error":"%s", "detail": "%s"}`, err.Message, err.Error), err.Code)
		return
	}
	if response == nil {
		log.Printf("ERROR: response from method is nil\n")
		http.Error(w, "Internal server error. Check the logs.", http.StatusInternalServerError)
		return
	}

	// turn the response into JSON
	bytes, e := json.Marshal(response)
	if e != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	// send the response and log
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
	log.Printf("%s %s %s %d", r.RemoteAddr, r.Method, r.URL, 200)
}

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var defaultKeyFunc jwt.Keyfunc = func(*jwt.Token) (interface{}, error) {
			return privKey, nil
		}
		jsonWebTokenParsed, err := jwt.Parse(r.Header.Get("jwt"), defaultKeyFunc)
		if err != nil || !jsonWebTokenParsed.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
