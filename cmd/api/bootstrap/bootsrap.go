package bootstrap

import (
	"context"
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

	fetchingService := fetching.NewFetchingService(courseRepository)

	createService := creating.NewCreatingService(courseRepository)

	commandBus := inmemory.New()

	commandBus.Register(creating.CreateCourseCommandType, creating.NewCreateCommandHandler(createService))
	commandBus.Register(fetching.ListCoursesCommandType, fetching.NewListCommandHandler(fetchingService))

	ctx, srv := server.New(context.Background(), host, port, commandBus)
	return srv.Run(ctx)
}
