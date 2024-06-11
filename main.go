// main.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.LoadHTMLGlob("html-files/*")

    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "about-us.html", gin.H{
            "title": "Welcome to Bark & Beyond!",
            "content": "Welcome to Bark & Beyond!",
        })
    })

    r.GET("/about", func(c *gin.Context) {
        c.HTML(http.StatusOK, "about-us.html", gin.H{
            "title": "About Us",
            "content": "This is the About Us page content.",
        })
    })

    r.GET("/services", func(c *gin.Context) {
        c.HTML(http.StatusOK, "services.html", gin.H{
            "title": "Services",
            "content": "This is the Services page content.",
        })
    })

    r.GET("/reviews", func(c *gin.Context) {
        c.HTML(http.StatusOK, "reviews.html", gin.H{
            "title": "Reviews",
            "content": "This is the Reviews page content.",
        })
    })

    r.Run(":8080")
}
