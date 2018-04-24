package main

import (
	"log"
	. "net/http"
)

func Index(w ResponseWriter, r Request) {

}

func main() {
	err := ListenAndServe(":7000", Handle("/", HandleFunc(Index)))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
