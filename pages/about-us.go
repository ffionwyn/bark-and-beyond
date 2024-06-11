package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AboutUsHandler(c *gin.Context) {
    c.String(http.StatusOK, "about us page")
}
