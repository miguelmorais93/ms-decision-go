package router

import "github.com/gin-gonic/gin"

func InitializeRouterConfig() {

	// Initalize router
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	// Initalize server
	router.Run(":8080")
}
