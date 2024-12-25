package router

import (
	"ms-decision-go/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")

	{
		v1.GET("/decision", handler.DecisionHandler)
	}

	{
		v1.POST("/decision/create-decision", handler.CreateDecisionHandler)
	}

	{
		v1.DELETE("/decision/delete-decision", handler.DeleteDecisionHandler)
	}

	{
		v1.PUT("/decision/change-decision", handler.ChangeDecisionHandler)
	}

	{
		v1.GET("/decision/list", handler.ListDecisionHandler)
	}
}
