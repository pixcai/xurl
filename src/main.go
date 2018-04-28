package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const (
	ServerAddr = ":7000"
	RedisAddr  = "localhost:6379"
)

func main() {
	handler := mux.NewRouter()
	handler.HandleFunc("/", Index).Methods("GET", "POST")
	handler.HandleFunc("/{hash}", Redirect).Methods("GET")

	server := &http.Server{
		Handler:      handler,
		Addr:         ServerAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("xUrlServer is running on " + ServerAddr)
	log.Fatal(server.ListenAndServe())
}
