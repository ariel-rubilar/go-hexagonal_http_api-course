package server

import (
	"fmt"

	mooc "github.com/ariel-rubilar/go-hexagonal_http_api-course/internal"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/platform/server/handler/courses"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/platform/server/handler/health"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Run() error
}

type api struct {
	httpAddr string
	engine   *gin.Engine

	courseRepository mooc.CourseRepository
}

func New(host string, port int, cr mooc.CourseRepository) Server {
	api := &api{
		httpAddr:         fmt.Sprintf("%s:%d", host, port),
		engine:           gin.New(),
		courseRepository: cr,
	}

	api.registerRoutes()
	return api
}

func (a *api) Run() error {
	fmt.Println("Starting server on", a.httpAddr)
	return a.engine.Run(a.httpAddr)
}

func (a *api) registerRoutes() {

	a.engine.GET("/health", health.CheckHandler())
	a.engine.POST("/courses", courses.CreateHandler(a.courseRepository))
}
