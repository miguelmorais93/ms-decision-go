package handler

import (
	"database/sql"
	"ms-decision-go/config"
)

var (
	db     *sql.DB
	logger *config.Logger
)

func InitalizeHandler() {
	logger = config.GetLogger("handler: ")
	db = config.Getdatabase()

}
