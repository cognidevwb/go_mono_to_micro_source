package platform

import "sync"

// Handler consumes one event payload.
type Handler func(payload any)

// Bus is an in-process publish/subscribe bus. Splitting the monolith turns every
// topic here into a broker topic — nothing in the compiler says so.
type Bus struct {
	mu       sync.RWMutex
	handlers map[string][]Handler
}

// NewBus returns an empty bus.
func NewBus() *Bus { return &Bus{handlers: map[string][]Handler{}} }

// Subscribe registers h for topic.
func (b *Bus) Subscribe(topic string, h Handler) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.handlers[topic] = append(b.handlers[topic], h)
}

// Publish delivers payload to every subscriber of topic, synchronously.
func (b *Bus) Publish(topic string, payload any) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	for _, h := range b.handlers[topic] {
		h(payload)
	}
}
