package main

import (
	"api-golang-crud/database"
	"api-golang-crud/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	Init()
	println(os.Getenv("JWT_SECRET"))
	database.Connect()

	r := gin.Default()

	err := r.SetTrustedProxies([]string{"192.168.0.108"})
	if err != nil {
		panic(err)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, Go Gin Framework!"})
	})

	routes.AuthRoutes(r)
	routes.UserRoutes(r)

	r.Run(":8080") // Jalankan server di port 8080
}
