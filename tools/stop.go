package tools

import (
	"log"
	"os"
	"os/signal"
)

func StopApp() {

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Shutting down...")
}
