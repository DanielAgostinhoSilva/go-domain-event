package events

import (
	"context"
	"sync"
)

// MockEventHandler é uma implementação de EventHandler para testes
type MockEventHandler struct {
	Called chan bool
}

// NewMockEventHandler cria uma nova instância de MockEventHandler
func NewMockEventHandler() *MockEventHandler {
	return &MockEventHandler{
		Called: make(chan bool, 1), // Canal com buffer de 1 para evitar bloqueio em Handle
	}
}

// Handle define o método Handle do EventHandler para MockEventHandler
func (m *MockEventHandler) Handle(ctx context.Context, event Event, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Called <- true
}
