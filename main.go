package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hasyimzii/go-notes/routes"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	routes.ApiRoutes(api)

	r.Run("localhost:3000")
}
