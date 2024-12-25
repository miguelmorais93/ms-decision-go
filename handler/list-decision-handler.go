package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListDecisionHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "decision1",
	})
}
