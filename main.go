package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" 
    }

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

    r.GET("/book-meet-greet", func(c *gin.Context) {
        c.HTML(200, "book-meet-greet.html", gin.H{})
    })

    r.Run(":" + port)
}
