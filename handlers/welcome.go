package handlers

import (
	"net/http"
)

func WelcomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/welcome.html")
}
