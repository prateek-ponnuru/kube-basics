package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Grab time of startup
	started := time.Now()

	// Connect to database, or Crash
	if err := connectDatabase(); err != nil {
		log.Fatal(err)
	}
	defer databaseConn.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Calculate runtime duration
		duration := time.Now().Sub(started)
		if duration.Seconds() > 100 {
			log.Println("Timeout triggered")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`hello gopher`))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`hello gopher`))
		}
	})

	http.HandleFunc("/aligator", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi Mr Aligator")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
