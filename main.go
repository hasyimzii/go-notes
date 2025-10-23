package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hasyimzii/go-notes/routes"
)

func main() {
	r := gin.Default()

	routes.V1Routes(r)

	r.Run("localhost:3000")
}
