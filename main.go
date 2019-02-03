package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ultrasad/goapi/handlers"
)

func main() {
	log.Print("Starting the service")

	//Router with Net/Http
	//handlers.Router()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	//Router with Muxa
	router := handlers.RouterMux()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	log.Print("The service is ready to listen and serve.")
	//log.Fatal(http.ListenAndServe(":1323", router))
	//log.Fatal(http.ListenAndServe(":1323", nil))

	//log.Fatal(http.ListenAndServe(":"+port, router))

	killSignal := <-stop
	switch killSignal {
	case os.Interrupt:
		log.Print("Got SIGINT...")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM...")
	}

	log.Print("The service is shutting down...")
	srv.Shutdown(context.Background())
	log.Print("Server gracefully stopped")
}
