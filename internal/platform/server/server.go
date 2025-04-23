package server

import (
	"fmt"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/platform/server/handler/health"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Run() error
}

type api struct {
	host string
	port int
}

func New(host string, port int) Server {
	return &api{
		host: host,
		port: port,
	}
}

func (a *api) Run() error {
	httpAddr := fmt.Sprintf("%s:%d", a.host, a.port)

	fmt.Println("Starting server on", httpAddr)

	srv := gin.New()
	srv.GET("/health", health.Handler)

	return srv.Run(httpAddr)
}
