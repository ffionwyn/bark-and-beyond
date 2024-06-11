package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReviewsHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "reviews.html", gin.H{
        "title": "Reviews",
        "content": "This is the Reviews page content.",
    })
}
