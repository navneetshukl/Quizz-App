package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")

	router.GET("/", func(c *gin.Context) {
		/*c.JSON(http.StatusOK, gin.H{
			"message": "This is Home Page",
		})*/
		c.HTML(http.StatusOK, "home.html", gin.H{})
		//name := c.DefaultQuery("name", "")
		fmt.Println(c.Param("name"))
		

	})

	router.Run()
}
