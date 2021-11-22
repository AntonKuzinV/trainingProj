package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"time"
	"trainingProj/websockets"
)

func main() {
	topic := "quickstart-events"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader: ", err)
	}

	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	//writing some messages
	err = w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	err = conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	if err != nil {
		return
	}
	//read some messages
	batch := conn.ReadBatch(10e3, 1e6)

	b := make([]byte, 10e3)
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}
	if err = conn.Close(); err != nil {
		log.Fatal("failed to close connection: ", err)
	}
	http.Handle("/", websocket.Handler(websockets.Echo))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
