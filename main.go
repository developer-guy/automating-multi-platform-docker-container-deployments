package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const GreetingMessage = "Hello World Docker Istanbul Community Meetup Group!"

func main() {
	m := os.Getenv("GREETING_MESSAGE")
	if m == "" {
		m = GreetingMessage
	}
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	http.HandleFunc("/greeting", greeting(m))

	go func() {
		log.Fatalln(http.ListenAndServe(":8080", nil))
	}()

	log.Println("Server is ready to handle requests at :8080")

	<-ch

	log.Println("Shutting down...")
}

func greeting(m string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(m))
	}
}
