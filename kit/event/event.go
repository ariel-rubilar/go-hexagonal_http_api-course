package event

import (
	"context"
	"time"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/domain/uuid"
)

type Type string

type Event interface {
	ID() string
	AggregateID() string
	OccurredOn() time.Time
	Type() Type
}

type BaseEvent struct {
	id          string
	aggregateID string
	occurredOn  time.Time
}

func NewBaseEvent(aggregateID string) BaseEvent {
	return BaseEvent{
		id:          uuid.New().String(),
		aggregateID: aggregateID,
		occurredOn:  time.Now(),
	}
}

func (e BaseEvent) ID() string {
	return e.id
}

func (e BaseEvent) AggregateID() string {
	return e.aggregateID
}

func (e BaseEvent) OccurredOn() time.Time {
	return e.occurredOn
}

type Handler interface {
	Handle(context.Context, Event) error
}

type Bus interface {
	Publish(context.Context, []Event) error
	Subscribe(Type, Handler) error
}
