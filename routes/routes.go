package routes

import (
	"Quizz-App/database"
	"Quizz-App/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	result:=[]models.CATEGORIES{}
	db,_:=database.ConnectToDatabase()

	db.Find(&result)
	n:=len(result)
	data:=[]string{}
	for i:=0;i<n;i++ {
		data=append(data,result[i].Category)
	}
	
	c.HTML(http.StatusOK, "home.html", gin.H{
		"data": data,
	})

}

func PostHome(c *gin.Context) {
	option := c.PostForm("option")
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s", option))
}

func GetCategory(c *gin.Context) {
	option := c.Param("option")
	c.String(http.StatusOK, "Option selected: %s", option)
}

func GetQuestion(c *gin.Context) {
	c.HTML(200,"question.html",gin.H{})
	
}

func AddQuestions(c *gin.Context) {
	category := c.PostForm("category")
	question := c.PostForm("question")
	option1 := c.PostForm("option1")
	option2 := c.PostForm("option2")
	option3 := c.PostForm("option3")
	option4 := c.PostForm("option4")
	answer := c.PostForm("answer")

	db,err:=database.ConnectToDatabase()
	if err!=nil{
		log.Fatal("Unable to Connect to Database ",err)
		return
	}
	
	
	questions:=&models.QUESTIONS{
		Category:category,
		Question:question,
		Option1:option1,
		Option2:option2,
		Option3:option3,
		Option4:option4,
		Answer:answer,
	}
	cat:=&models.CATEGORIES{
		Category: category,
	}
	db.Create(&questions)
	db.Create(&cat)
	c.Redirect(http.StatusSeeOther,"/addquestions")
}
