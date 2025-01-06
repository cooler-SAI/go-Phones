package tools

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func StopApp() {
	go func() {
		log.Println("Server starting...")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Printf("Server error: %v\n", err)
		}
	}()

	site := "http://localhost:8080"
	fmt.Printf("Open %s in browser\n", site)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Shutting down...")
}
