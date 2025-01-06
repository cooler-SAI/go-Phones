package handlers

import (
	"fmt"
	"net/http"
)

func WelcomePage(w http.ResponseWriter, r *http.Request) {
	fprintf, err := fmt.Fprintf(w, "Hello! Welcome to the Phone Site!")
	if err != nil {
		return
	}
	fmt.Println(w, fprintf)
}
