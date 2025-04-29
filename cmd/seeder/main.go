package main

import (
	"core.yuyuid.id/internal/database"
	"core.yuyuid.id/internal/database/seeders"
	"core.yuyuid.id/pkg/config"
)

func main() {
	cnf := config.Get()
	db := database.Connect(cnf.Database)

	seeders.Run(db)
}
