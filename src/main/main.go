package main

import (
	"chatServer/src/secondary"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Login function which redirects to html page home.html if connection is made
func Login(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home.html",
		gin.H{
			"title": "Home Page",
		},
	)
}

func main() {
	router := gin.Default()                       //creating gin router using default middleware
	router.LoadHTMLGlob("web/template/home.html") //The html webpage to load
	router.Static("/css", "../web/css")
	router.GET("/login", Login) //GET method called using /login and uses function Login
	router.POST("/sub", func(c *gin.Context) {
		router.LoadHTMLGlob("web/template/chatpage.html") //Loading html page chatpage
		name := c.PostForm("uname")
		secondary.Chatupdate()
		c.HTML( //To open html page instead of just json data
			http.StatusOK,
			"chatpage.html",
			gin.H{
				"title": "Chat Page",
				"uname": name,
			})
			c.JSON(
				http.StatusOK,
				gin.H{
					"title": "Chat Page",
					"uname": name,
				})
	})
	router.POST("/loaddata", func(c *gin.Context) {

	})
	router.Run(":8080") //Running router on port 0.0.0.0:8080
}
