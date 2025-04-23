package bootstrap

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const httpAddr = ":8080"

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func Run() error {
	fmt.Println("Starting server on", httpAddr)

	srv := gin.New()
	srv.GET("/health", healthHandler)

	log.Fatal(srv.Run(httpAddr))

	return nil
}
