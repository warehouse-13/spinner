package handlers

import (
	"log"
	"net/http"
)

func Spin(http.ResponseWriter, *http.Request) {
	log.Println("starting")
	for {
		for i := 0; i < 100000000; i++ {
		}
	}
}

func Unspin(http.ResponseWriter, *http.Request) {
	log.Println("stopping")
}
