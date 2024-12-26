package handler

import (
	"ms-decision-go/handler/requests"

	"github.com/gin-gonic/gin"
)

func CreateDecisionHandler(ctx *gin.Context) {

	var requestFrom requests.CreateDecisionRequest

	ctx.BindJSON(&requestFrom)
	logger.Infof("request : %v", requestFrom)
}
