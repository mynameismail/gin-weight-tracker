package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WeightController struct {}

func (h WeightController) Create(c *gin.Context) {
	c.HTML(http.StatusOK, "weight_create.html", gin.H{})
}

func (h WeightController) Store(c *gin.Context) {
	date := c.PostForm("date")
	weight := c.PostForm("weight")

	if date == "" || weight == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	c.Redirect(http.StatusFound, "/")
	return
}
