package handler

import "github.com/gin-gonic/gin"

func Authorization(c *gin.Context) {

}

func Registration(c *gin.Context) {

}

func Home(c *gin.Context) {
	c.HTML(200, "home.html", nil)
}

func Contact(c *gin.Context) {
	c.HTML(200, "contact.html", nil)
}
