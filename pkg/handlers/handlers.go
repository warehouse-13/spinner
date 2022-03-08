package handlers

import (
	"log"
	"net/http"
)

var quit chan struct{}

func StartSpinning() {
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
				for i := 0; i < 1000; i++ {
				}
			}
		}
		log.Println("stopped spinning")
	}()
}

func Spin(http.ResponseWriter, *http.Request) {
	StartSpinning()
}

func Unspin(http.ResponseWriter, *http.Request) {
	log.Println("sending signal to stop")
	if quit != nil {
		close(quit)
	} else {
		log.Println("not spinning yet")
	}
}
