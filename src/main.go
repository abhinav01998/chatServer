package main

import(
    "github.com/gin-gonic/gin"
    "net/http"
    )

func Login(c *gin.Context){
  c.HTML(
    http.StatusOK,
    "home.html",
    gin.H{
      "title": "Home Page",
      },
  )
}

func main(){
  router := gin.Default()
  router.LoadHTMLGlob("web/home.html")
  router.Static("/css", "../web/css")
  router.GET("/login",Login)
  router.POST("/sub", func(c *gin.Context) {
    router.LoadHTMLGlob("web/chatpage.html")
      name := c.PostForm("uname")
        c.HTML(
            http.StatusOK,
            "chatpage.html",
            gin.H{
              "title": "Chat Page",
              "uname": name,
            })
        })
	router.Run(":8080")
}
