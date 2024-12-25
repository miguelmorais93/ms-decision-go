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

func DeleteDecisionHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "decision1",
	})
}

func ChangeDecisionHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "decision1",
	})
}

func ListDecisionHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "decision1",
	})
}

func DecisionHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "decision1",
	})
}
