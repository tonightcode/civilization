package main

import (
	. "culture/routes"
	"log"
)

// go run github.com/cosmtrek/air
func main() {
	startWebServer(":8080")
}

func startWebServer(port string) {
	r := NewRouter()
	log.Println("Starting HTTP service at " + port)
	r.Run(port)
}
