package bootstrap

import (
	"fmt"

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

	srv := server.New(host, port, courseRepository)
	return srv.Run()
}
