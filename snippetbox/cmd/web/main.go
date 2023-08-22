package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
)


func main() {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
