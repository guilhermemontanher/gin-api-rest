package main

import (
	"github.com/guilhermemontanher/api-go-gin/database"
	"github.com/guilhermemontanher/api-go-gin/routes"
)

func main() {
	database.ConnectDatabase()
	routes.HandleRequests()
}
