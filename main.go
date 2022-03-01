package main

import (
	"net/http"

	"github.com/callisto13/spinner/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/spin", handlers.Spin)
	r.HandleFunc("/unspin", handlers.Unspin)

	_ = http.ListenAndServe(":8080", r)
}
