package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hasyimzii/go-notes/handlers"
)

func ApiRoutes(r *gin.RouterGroup) {
	// Notes
	noteGroup := r.Group("/notes")
	{
		noteGroup.GET("/", handlers.GetNotes)
		noteGroup.POST("/", handlers.GetNotes)
		noteGroup.PUT("/", handlers.GetNotes)
		noteGroup.DELETE("/", handlers.GetNotes)
	}
}
