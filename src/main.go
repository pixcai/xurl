package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ServerAddr := os.Getenv("XURL_URL")
	if len(ServerAddr) == 0 {
		ServerAddr = ":9900"
	}

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
