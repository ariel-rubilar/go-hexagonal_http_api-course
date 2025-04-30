package server

import (
	"context"
	"fmt"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/platform/server/handler/courses"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/platform/server/handler/health"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/kit/command"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Run(context.Context) error
}

type api struct {
	httpAddr string
	engine   *gin.Engine

	bus command.Bus
}

func New(ctx context.Context, host string, port int, b command.Bus) (context.Context, Server) {
	api := &api{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.New(),
		bus:      b,
	}

	api.registerRoutes()

	return ctx, api
}

func (a *api) Run(ctx context.Context) error {
	fmt.Println("Starting server on", a.httpAddr)
	return a.engine.Run(a.httpAddr)
}

func (a *api) registerRoutes() {

	a.engine.GET("/health", health.CheckHandler())
	a.engine.POST("/courses", courses.CreateHandler(a.bus))
	a.engine.GET("/courses", courses.ListHandler(a.bus))
}
