package service

import (
	"fmt"
	"net/http"
)

const (
	WardenStampHeader = "X-Auth-Authorized"
	UserNameHeaer     = "X-Auth-Username"
)

func WardenStampChecker(h http.Handler) http.Handler {
	fmt.Println("Middleware WardenStampChecker initialized....")
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware WardenStampChecker func called....")
		authFlag := r.Header.Get(WardenStampHeader)
		if authFlag != "true" {
			http.Error(w, "warden auth stamp not found", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func UserNameChecker(h http.Handler) http.Handler {
	fmt.Println("Middleware UserNameChecker initialized....")
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware UserNameChecker func called....")
		fmt.Println(r.Header)
		userName := r.Header.Get(UserNameHeaer)
		if userName == "" {
			http.Error(w, "userName not found", http.StatusBadRequest)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
