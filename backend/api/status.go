package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Status is instance of Controller and provides status to liveness/readiness probes.
//
// Parameters:
// - A pointer to gin.Context
//
// Returns:
// - N/A
func (c *Controller) Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
