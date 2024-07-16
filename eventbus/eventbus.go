package eventbus

import (
	"sync"
)

type EventBus struct {
	subscribers map[string][]chan any
	lock        sync.RWMutex
}

var Bus *EventBus

func init() {
	Bus = &EventBus{
		subscribers: make(map[string][]chan any),
	}
}

func (eb *EventBus) Subscribe(event string, listener chan any) {
	eb.lock.Lock()
	defer eb.lock.Unlock()

	if _, ok := eb.subscribers[event]; !ok {
		eb.subscribers[event] = make([]chan any, 0)
	}
	eb.subscribers[event] = append(eb.subscribers[event], listener)
}

func (eb *EventBus) Unsubscribe(event string, listener chan any) {
	eb.lock.Lock()
	defer eb.lock.Unlock()

	if subscribers, ok := eb.subscribers[event]; ok {
		for i, ch := range subscribers {
			if ch == listener {
				eb.subscribers[event] = append(subscribers[:i], subscribers[i+1:]...)
				close(listener)
				break
			}
		}
	}
}

func (eb *EventBus) Publish(event string, data any) {
	subscribers, ok := eb.subscribers[event]
	if !ok {
		return
	}

	for _, ch := range subscribers {
		ch <- data
	}
}
