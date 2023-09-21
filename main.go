package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jrandyl/webapp/web"
)

func main() {
	r := gin.Default()

	web.Static(r)

	r.GET("/api", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello")
	})
	// Start the server
	err := r.Run(":4000")
	if err != nil {
		log.Fatalf("could not start the server: %v", err)
	}
}
