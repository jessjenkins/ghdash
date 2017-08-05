package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting server")

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/hello", homeHandler)
	r.HandleFunc("/{repoName}/", repoHandler)

	defer log.Println("Stopping Server")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello World"))
}

func repoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	repoName := vars["repoName"]
	log.Printf("Serving repo:%s\n", repoName)

	w.Write([]byte(fmt.Sprintf("Hello %s!\n", repoName)))
}
