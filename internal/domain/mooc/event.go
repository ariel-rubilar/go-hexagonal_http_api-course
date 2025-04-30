package mooc

import "github.com/ariel-rubilar/go-hexagonal_http_api-course/kit/event"

var (
	CreatedCourseEventType = event.Type("event.created.course")
)

type CreatedEvent struct {
	id       string
	name     string
	duration string
	event.BaseEvent
}

func NewCreatedEvent(id, name, duration string) *CreatedEvent {
	return &CreatedEvent{
		id:        id,
		name:      name,
		duration:  duration,
		BaseEvent: event.NewBaseEvent(id),
	}
}

func (e *CreatedEvent) Type() event.Type {
	return CreatedCourseEventType
}

func (e *CreatedEvent) ID() string {
	return e.id
}

func (e *CreatedEvent) Name() string {
	return e.name
}

func (e *CreatedEvent) Duration() string {
	return e.duration
}
