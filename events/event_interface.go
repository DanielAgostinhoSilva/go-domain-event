package events

import (
	"context"
	"sync"
	"time"
)

type Event interface {
	GetID() string
	GetType() string
	GetAggregateID() string
	GetAggregateType() string
	GetTimestamp() time.Time
	GetVersion() int
	GetData() interface{}
	GetMetadata() map[string]string
}

type EventHandler interface {
	Handle(ctx context.Context, event Event, wg *sync.WaitGroup)
}

type EventDispatcher interface {
	Register(eventName string, handler EventHandler) error
	Dispatch(ctx context.Context, event Event) error
	Remove(eventName string, handler EventHandler) error
	Has(eventName string, handler EventHandler) bool
	Clear()
}
