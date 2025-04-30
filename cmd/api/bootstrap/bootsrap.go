package bootstrap

import (
	"fmt"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/application/course/creating"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/application/course/fetching"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/platform/bus/inmemory"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/platform/persistence/memdb"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/platform/persistence/memdb/course"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/platform/server"
)

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	db, err := memdb.NewMemDB()
	if err != nil {
		return fmt.Errorf("error creating memsql: %w", err)
	}
	courseRepository := course.NewCourseRepository(db)

	courseService := fetching.NewCourseService(courseRepository)

	createService := creating.NewCourseService(courseRepository)

	commandBus := inmemory.New()

	commandBus.Register(creating.CreateCourseCommandType, creating.NewCreateCommandHandler(createService))
	commandBus.Register(fetching.ListCoursesCommandType, fetching.NewListCommandHandler(courseService))

	srv := server.New(host, port, commandBus)
	return srv.Run()
}
