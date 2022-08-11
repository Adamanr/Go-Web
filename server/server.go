package Server

import (
	"GoWeb/server/handler"
	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	r.LoadHTMLGlob("./client/templates/html/*")

	users := r.Group("/users")
	users.POST("/auth", handler.Authorization)
	users.POST("/reg", handler.Registration)

	page := r.Group("/page")
	page.GET("/", handler.Home)
	page.GET("/contact", handler.Contact)

	r.Run(":8080")
}
