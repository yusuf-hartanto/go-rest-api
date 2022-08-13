package main

import (
	"rest-api/config"
)

func main() {
	db, _ := config.DBConnection()
	config.SetupRoutes(db)
}
