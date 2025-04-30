package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	return serverContext(ctx), api
}

func (a *api) Run(ctx context.Context) error {
	fmt.Println("Starting server on", a.httpAddr)

	srv := &http.Server{
		Addr:    a.httpAddr,
		Handler: a.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return srv.Shutdown(ctx)
}

func (a *api) registerRoutes() {

	a.engine.GET("/health", health.CheckHandler())
	a.engine.POST("/courses", courses.CreateHandler(a.bus))
	a.engine.GET("/courses", courses.ListHandler(a.bus))
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		<-c
		cancel()
	}()

	return ctx
}
