package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	data:=[]string{"option1","option2","option3"}
	c.HTML(http.StatusOK, "home.html", gin.H{
		"data":data,
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
