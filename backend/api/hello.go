package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Status is instance of Controller and provides demo Hello message.
//
// Parameters:
// - A pointer to gin.Context
//
// Returns:
// - N/A
func (c *Controller) Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello Everybody! How are you today?",
	})
}
