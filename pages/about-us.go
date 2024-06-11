package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AboutUsHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "about-us.html", gin.H{
        "title": "About Us",
        "content": "This is the About Us page content.",
    })
}
