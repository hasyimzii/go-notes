package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNotes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
