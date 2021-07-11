package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8080"

	router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and running...")
	})
	router.HandleFunc("/post", getPosts).Methods("GET")
	router.HandleFunc("/post", addPost).Methods("POST")
	log.Println("Server: listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
