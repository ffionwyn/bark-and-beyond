package pages

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func ServicesHandler(c *gin.Context) {
    c.String(http.StatusOK, "services page")
}
