package main

import (
	"log"
	"net/http"

	"github.com/callisto13/spinner/pkg/handlers"
	"github.com/gorilla/mux"
)

const port = ":8080"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/spin", handlers.Spin)
	r.HandleFunc("/unspin", handlers.Unspin)

	log.Printf("listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
