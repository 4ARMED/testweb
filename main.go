package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	message := "Pong. You hit me with a %v request\n" + html.EscapeString(r.Method)

	w.Write([]byte(message))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", echoHandler)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	bindAddress := flag.String("b", "127.0.0.1:8000", "Bind address in format IP:PORT")
	flag.Parse()

	srv := &http.Server{
		Handler:      loggedRouter,
		Addr:         *bindAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Starting web server on %v\n", *bindAddress)
	log.Fatal(srv.ListenAndServe())
}
