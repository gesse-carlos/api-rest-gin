package main

import (
	"github.com/gesse-carlos/gin-api-rest/database"
	"github.com/gesse-carlos/gin-api-rest/routes"
)

func main() {
	database.Connection()
	routes.HandleRequests()
}
