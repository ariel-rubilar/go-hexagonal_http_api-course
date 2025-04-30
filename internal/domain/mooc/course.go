package mooc

import "github.com/ariel-rubilar/go-hexagonal_http_api-course/kit/event"

type Course struct {
	id       CourseID
	name     CourseName
	duration CourseDuration
	events   []event.Event
}

func NewCourse(id, name, duration string) (*Course, error) {

	courseID, err := NewCourseID(id)

	if err != nil {
		return nil, err
	}

	courseName, err := NewCourseName(name)
	if err != nil {
		return nil, err
	}

	courseDuration, err := NewCourseDuration(duration)

	if err != nil {
		return nil, err
	}

	course := &Course{
		id:       courseID,
		name:     courseName,
		duration: courseDuration,
	}

	course.RecordEvent(NewCreatedEvent(courseID.String(), courseName.String(), courseDuration.String()))

	return course, nil
}

func (c *Course) ID() CourseID {
	return c.id
}

func (c *Course) Name() CourseName {
	return c.name
}

func (c *Course) Duration() CourseDuration {
	return c.duration
}

func (c *Course) PullEvents() []event.Event {
	events := c.events
	c.events = []event.Event{}
	return events
}

func (c *Course) RecordEvent(e event.Event) {
	c.events = append(c.events, e)
}
