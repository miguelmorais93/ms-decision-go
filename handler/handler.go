package handler

import (
	"ms-decision-go/config"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *config.Logger
)

func InitalizeHandler() {
	logger = config.GetLogger("handler: ")
	db = config.Getdatabase()

}
