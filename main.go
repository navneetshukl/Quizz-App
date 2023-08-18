package main

import (
	"Quizz-App/database"
	"Quizz-App/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")
	database.MigrateDatabase()

	router.GET("/", routes.GetHome)

	router.POST("/", routes.PostHome)

	router.GET("/:option", routes.GetCategory)

	router.Run(":8080")
}
