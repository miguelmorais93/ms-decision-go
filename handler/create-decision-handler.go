package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDecisionHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "decision1",
	})
}