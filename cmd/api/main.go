package main

import (
	"fmt"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		panic(fmt.Errorf("error starting the application: %v", err))
	}
}
