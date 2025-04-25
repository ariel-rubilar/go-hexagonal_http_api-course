package server

import (
	"fmt"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/application/course"
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

	courseService course.CourseService
}

func New(host string, port int, s course.CourseService) Server {
	api := &api{
		httpAddr:      fmt.Sprintf("%s:%d", host, port),
		engine:        gin.New(),
		courseService: s,
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
	a.engine.POST("/courses", courses.CreateHandler(a.courseService))
	a.engine.GET("/courses", courses.ListHandler(a.courseService))
}
