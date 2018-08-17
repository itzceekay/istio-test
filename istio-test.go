package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func regroute() {
	port := 8001
	router = mux.NewRouter()
	router.HandleFunc("/here", mainEmail).Methods("GET")
	address := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(address, router))
}

func mainEmail(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	response := "You are here"
	json.NewEncoder(w).Encode(response)

}

func main() {
	go regroute()
	forever := make(chan bool)
	<-forever
}
