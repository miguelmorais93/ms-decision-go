package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteDecisionHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "decision1",
	})
}
