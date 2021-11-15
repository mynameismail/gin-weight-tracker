package routes

import (

	"net/http"

	"github.com/gin-gonic/gin"
)

func MakeAppRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello!")
	})
}
