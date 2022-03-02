package handlers

import (
	"log"
	"net/http"
)

var quit chan struct{}

func Spin(http.ResponseWriter, *http.Request) {
	quit = (make(chan struct{}))
	log.Println("starting to spin")
	go func() {
	loop:
		for {
			select {
			case <-quit:
				log.Println("received signal to stop")
				break loop
			default:
				for i := 0; i < 3; i++ {
				}
			}
		}
	}()

	log.Println("stopped")
}

func Unspin(http.ResponseWriter, *http.Request) {
	log.Println("sending signal to stop")
	close(quit)
}
