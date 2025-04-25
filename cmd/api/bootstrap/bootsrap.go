package bootstrap

import (
	"fmt"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/application/course"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/platform/persistence/memdb"
	courseRepo "github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/platform/persistence/memdb/course"
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
	courseRepository := courseRepo.NewCourseRepository(db)

	courseService := course.NewCourseService(courseRepository)

	srv := server.New(host, port, courseService)
	return srv.Run()
}
