package main

import (
	"demo/handlers"
	"log/slog"
	"net/http"
	"runtime"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/health", Health)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/user", handlers.CreateUser) // create user

	slog.Info("The application is listening on port 8083")
	if err := http.ListenAndServe("0.0.0.0:8083", nil); err != nil {
		slog.Info("Application has returned error", err)
		runtime.Goexit()
	}

	// all network interfaces:8083
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
	w.WriteHeader(http.StatusOK)
}

// Cloud Native Applications
// Go does not require a webserver like Apache Tomcat/IIS
// It has self hosted applications
// The binary itself contains a packed webserver, the moment it see http package

// listener
// Router
// ResponseWriter
// Request <- All the request data , body, headers, params, authentication heardss
