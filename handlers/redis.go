package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"main/clients"
	"main/utils"
	"net/http"

	"github.com/brownhounds/swift/res"
)

func PublishEventHandlerRedis(w http.ResponseWriter, r *http.Request) {
	token := utils.GenerateSecureToken(10)
	client := clients.Redis.Client()

	if err := client.Publish(r.Context(), "events", token).Err(); err != nil {
		panic(err)
	}

	res.Json(w, http.StatusOK, res.Map{
		"message_published": token,
	})
}

func SseHandlerRedis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("Streaming Unsupported")
	}

	client := clients.Redis.Client()
	sub := client.Subscribe(r.Context(), "events")

	go func() {
		<-r.Context().Done()
		log.Println("Client disconnected, cleaning up...")
		if err := sub.Unsubscribe(r.Context()); err != nil {
			panic(err)
		}
		sub.Close()
	}()

	for msg := range sub.Channel() {
		data, err := json.Marshal(res.Map{"message": msg.Payload})
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "id: %s\n", msg.Channel)
		fmt.Fprintf(w, "event: %s\n", "test-event")
		fmt.Fprintf(w, "data: %s\n\n", data)
		flusher.Flush()
	}
}
