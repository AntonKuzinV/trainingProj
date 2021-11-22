package main

import (
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"trainingProj/websockets"
)

func main() {
	http.Handle("/", websocket.Handler(websockets.Echo))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
