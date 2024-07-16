package main

import (
	"main/handlers"

	"github.com/brownhounds/swift"
)

func main() {
	app := swift.New()

	app.Get("/redis-events", handlers.SseHandlerRedis)
	app.Get("/redis-publish", handlers.PublishEventHandlerRedis)

	app.Get("/bus-events", handlers.SseHandlerBus)
	app.Get("/bus-publish", handlers.PublishEventHandlerBus)

	app.RootStaticServer("static", false)

	app.Serve("0.0.0.0", "8080")
}
