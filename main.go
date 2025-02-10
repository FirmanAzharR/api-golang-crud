package main

import (
	"api-golang-crud/database"
	"api-golang-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()
	routes.UserRoutes(r)

	r.Run(":8080") // Jalankan server di port 8080
}
