package webserver

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

// Start ...
func Start(ctx context.Context) {
	r := gin.Default()

    registerRoutes(r)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
