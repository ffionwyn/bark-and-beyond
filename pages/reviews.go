package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReviewsHandler(c *gin.Context) {
    c.String(http.StatusOK, "reviews page")
}
