package events

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDomainEventDispatcher_Register(t *testing.T) {
	dispatcher := NewEventDispatcher()
	handler := NewMockEventHandler()

	err := dispatcher.Register("TestEvent", handler)
	require.NoError(t, err)

	// Verifica se o handler foi registrado
	assert.True(t, dispatcher.Has("TestEvent", handler))
}

func TestDomainEventDispatcher_Register_Duplicate(t *testing.T) {
	dispatcher := NewEventDispatcher()
	handler := NewMockEventHandler()

	err := dispatcher.Register("TestEvent", handler)
	require.NoError(t, err)

	// Tentar registrar o mesmo handler novamente deve resultar em erro
	err = dispatcher.Register("TestEvent", handler)
	assert.Equal(t, ErrHandlerAlreadyRegistered, err)
}

func TestDomainEventDispatcher_Dispatch(t *testing.T) {
	dispatcher := NewEventDispatcher()
	handler1 := NewMockEventHandler()
	handler2 := NewMockEventHandler()

	dispatcher.Register("TestEvent", handler1)
	dispatcher.Register("TestEvent", handler2)

	event := &BaseEvent{
		ID:            "1",
		Type:          "TestEvent",
		AggregateID:   "agg-1",
		AggregateType: "Type",
		Timestamp:     time.Now(),
		Version:       1,
		Data:          nil,
		Metadata:      map[string]string{},
	}

	// Despacha o evento
	dispatcher.Dispatch(context.Background(), event)

	// Verifica se os handlers foram chamados
	assert.True(t, <-handler1.Called)
	assert.True(t, <-handler2.Called)
}

func TestDomainEventDispatcher_Remove(t *testing.T) {
	dispatcher := NewEventDispatcher()
	handler := NewMockEventHandler()

	dispatcher.Register("TestEvent", handler)
	assert.True(t, dispatcher.Has("TestEvent", handler))

	// Remove o handler
	dispatcher.Remove("TestEvent", handler)
	assert.False(t, dispatcher.Has("TestEvent", handler))
}

func TestDomainEventDispatcher_Clear(t *testing.T) {
	dispatcher := NewEventDispatcher()
	handler := NewMockEventHandler()

	dispatcher.Register("TestEvent", handler)
	require.True(t, dispatcher.Has("TestEvent", handler))

	// Limpa todos os handlers
	dispatcher.Clear()
	assert.False(t, dispatcher.Has("TestEvent", handler))
}
