package events

import "time"

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
