package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.Static("/images", "./images")

    r.LoadHTMLGlob("html-files/*")

    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "homepage.html", gin.H{})
    })

    r.Run(":8080")
}
