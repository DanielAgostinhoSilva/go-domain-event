package events

import (
	"context"
	"errors"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type DomainEventDispatcher struct {
	handlers map[string][]EventHandler
	mu       sync.RWMutex
}

// NewEventDispatcher cria uma nova instância de DomainEventDispatcher
func NewEventDispatcher() *DomainEventDispatcher {
	return &DomainEventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

// Register registra um handler para um evento específico
func (ed *DomainEventDispatcher) Register(eventName string, handler EventHandler) error {
	ed.mu.Lock()
	defer ed.mu.Unlock()

	if handler == nil {
		return errors.New("handler cannot be nil")
	}

	if handlers, exists := ed.handlers[eventName]; exists {
		for _, h := range handlers {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

// Dispatch despacha um evento para todos os handlers registrados
func (ed *DomainEventDispatcher) Dispatch(ctx context.Context, event Event) error {
	ed.mu.RLock()
	defer ed.mu.RUnlock()

	if event == nil {
		return errors.New("event cannot be nil")
	}

	eventName := event.GetType()
	handlers, ok := ed.handlers[eventName]
	if !ok {
		return nil // Se não houver handlers registrados, não faz nada
	}

	var wg sync.WaitGroup
	for _, handler := range handlers {
		wg.Add(1)
		go handler.Handle(ctx, event, &wg)
	}
	wg.Wait()

	return nil
}

// Remove um handler específico de um evento
func (ed *DomainEventDispatcher) Remove(eventName string, handler EventHandler) error {
	ed.mu.Lock()
	defer ed.mu.Unlock()

	if handler == nil {
		return errors.New("handler cannot be nil")
	}

	handlers, ok := ed.handlers[eventName]
	if !ok {
		return nil // Se não houver handlers registrados, não faz nada
	}

	for i, h := range handlers {
		if h == handler {
			ed.handlers[eventName] = append(handlers[:i], handlers[i+1:]...)
			break
		}
	}
	return nil
}

// Has verifica se um handler está registrado para um evento específico
func (ed *DomainEventDispatcher) Has(eventName string, handler EventHandler) bool {
	ed.mu.RLock()
	defer ed.mu.RUnlock()

	if handler == nil {
		return false
	}

	handlers, ok := ed.handlers[eventName]
	if !ok {
		return false
	}

	for _, h := range handlers {
		if h == handler {
			return true
		}
	}
	return false
}

// Clear remove todos os handlers registrados
func (ed *DomainEventDispatcher) Clear() {
	ed.mu.Lock()
	defer ed.mu.Unlock()

	ed.handlers = make(map[string][]EventHandler)
}
