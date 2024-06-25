package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServicesHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "services.html", gin.H{
        "title": "Services",
        "content": "This is the Services page content.",
    })
}
