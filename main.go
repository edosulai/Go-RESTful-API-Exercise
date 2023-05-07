package main

import (
	"gorest/database"
	"gorest/router"
)

func main() {
	database.StartDB()

	router.New().Run(":3000")
}
