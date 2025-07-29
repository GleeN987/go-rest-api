package http

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
)

func JWTAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			http.Error(w, "not authorized1", http.StatusUnauthorized)
			return
		}

		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			http.Error(w, "not authorized2", http.StatusUnauthorized)
			return
		}

		if validateToken(authHeaderParts[1]) {
			original(w, r)
		} else {
			http.Error(w, "not authorized3", http.StatusUnauthorized)
			return
		}
	}
}

// eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiaWF0IjoxNzUzNjk0MjA1LCJuYW1lIjoiSm9obiBEb2UiLCJzdWIiOiIxMjM0NTY3ODkwIn0.ukaw5zrKGtEknrfDtNAQmgpuznoPzP-xcTSZub1OOF4vv99E5S8w04SlC3YGSrmHL6zYwyG-sYJqt4ostNQFOg
func validateToken(accessToken string) bool {
	var signingKey = []byte("gorestapikey")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error validating the token")
		}
		return signingKey, nil
	})

	if err != nil {
		fmt.Println(err)
		return false
	}

	return token.Valid
}
