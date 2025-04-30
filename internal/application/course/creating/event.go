package creating

import (
	"context"
	"fmt"
	"log"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/domain/mooc"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/kit/event"
)

type LogOnCourseCreated struct {
}

func NewLogOnCourseCreated() *LogOnCourseCreated {
	return &LogOnCourseCreated{}
}

func (h *LogOnCourseCreated) Handle(ctx context.Context, event event.Event) error {
	// Log the course creation event
	createdEvent, ok := event.(*mooc.CreatedEvent)

	if !ok {
		return fmt.Errorf("event is not of type CreatedEvent")
	}
	log.Printf("Course created: ID=%s, Name=%s, Duration=%s", createdEvent.ID(), createdEvent.Name(), createdEvent.Duration())

	return nil
}
