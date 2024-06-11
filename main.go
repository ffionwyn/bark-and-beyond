// main.go

package main

import (
	"net/http"

	"github.com/ffiongriffiths/bark-and-beyond/pages"
	"github.com/gin-gonic/gin"
)


func main() {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Welcome to Bark & Beyond!")
    })

	// routes
    r.GET("/about", pages.AboutUsHandler)
    r.GET("/services", pages.ServicesHandler)
    r.GET("/reviews", pages.ReviewsHandler)

    r.Run(":8080")
}
