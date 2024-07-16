package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"main/eventbus"
	"main/utils"
	"net/http"

	"github.com/brownhounds/swift/res"
)

func PublishEventHandlerBus(w http.ResponseWriter, r *http.Request) {
	token := utils.GenerateSecureToken(10)
	eventbus.Bus.Publish("test:event", token)

	res.Json(w, http.StatusOK, res.Map{
		"message_published": token,
	})
}

func SseHandlerBus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("Streaming Unsupported")
	}

	listener := make(chan any)

	eventbus.Bus.Subscribe("test:event", listener)

	go func() {
		<-r.Context().Done()
		log.Println("Client disconnected, cleaning up...")
		eventbus.Bus.Unsubscribe("test:event", listener)
	}()

	for event := range listener {
		data, err := json.Marshal(res.Map{"message": event})
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "id: %s\n", "something")
		fmt.Fprintf(w, "event: %s\n", "test-event")
		fmt.Fprintf(w, "data: %s\n\n", data)
		flusher.Flush()
	}
}
