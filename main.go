package main

import (
	"log"
	"net/http"
	"time"
)

func customHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func main() {
	port := "8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", customHandler)
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("./"))))

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger := log.Default()
	logger.Println("Starting server on port 8080...")
	logger.Fatal(server.ListenAndServe())
}
