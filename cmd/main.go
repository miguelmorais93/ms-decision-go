package main

import (
	"ms-decision-go/config"
	"ms-decision-go/router"
)

func main() {

	// Initialize configuration (dabatase, logger, ...)
	config.InitConfig()

	// Intialize router
	router.InitializeRouterConfig()

}
