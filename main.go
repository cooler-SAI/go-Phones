package main

import (
	"fmt"
	"go-Phones/handlers"
	"go-Phones/tools"
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server Starting....")

	http.HandleFunc("/", handlers.WelcomePage)

	log.Println("Server StopApp Running....")
	tools.StopApp()

	log.Println("Main End..")
}
