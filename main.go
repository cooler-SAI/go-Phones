package main

import (
	"fmt"
	"go-Phones/handlers"
	"go-Phones/tools"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server Starting....")

	http.HandleFunc("/", handlers.WelcomePage)

	log.Println("Server StopApp Running....")
	tools.StopApp()

	log.Println("Main End..")
}
