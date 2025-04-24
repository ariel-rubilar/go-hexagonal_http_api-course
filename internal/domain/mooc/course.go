package mooc

type Course struct {
	id       CourseID
	name     CourseName
	duration CourseDuration
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

	return &Course{
		id:       courseID,
		name:     courseName,
		duration: courseDuration,
	}, nil
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
