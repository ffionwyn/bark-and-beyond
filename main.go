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

    r.GET("/contact", func(c *gin.Context) {
        c.HTML(200, "contact.html", gin.H{})
    })

    r.GET("/about", func(c *gin.Context) {
        c.HTML(200, "about-us.html", gin.H{})
    })

    r.GET("/services", func(c *gin.Context) {
        c.HTML(200, "services.html", gin.H{})
    })

    r.GET("/reviews", func(c *gin.Context) {
        c.HTML(200, "reviews.html", gin.H{})
    })

    r.Run(":8080")
}
