package events

import "time"

type BaseEvent struct {
	ID            string
	Type          string
	AggregateID   string
	AggregateType string
	Timestamp     time.Time
	Version       int
	Data          interface{}
	Metadata      map[string]string
}

func (e BaseEvent) GetID() string {
	return e.ID
}
func (e BaseEvent) GetType() string {
	return e.Type
}
func (e BaseEvent) GetAggregateID() string {
	return e.AggregateID
}
func (e BaseEvent) GetAggregateType() string {
	return e.AggregateType
}
func (e BaseEvent) GetTimestamp() time.Time {
	return e.Timestamp
}
func (e BaseEvent) GetVersion() int {
	return e.Version
}
func (e BaseEvent) GetData() interface{} {
	return e.Data
}
func (e BaseEvent) GetMetadata() map[string]string {
	return e.Metadata
}
