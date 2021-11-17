package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {}

func (h ProfileController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "profile_index.html", gin.H{})
}

func (h ProfileController) Create(c *gin.Context) {
	c.HTML(http.StatusOK, "profile_create.html", gin.H{})
}

func (h ProfileController) Store(c *gin.Context) {
	username := c.PostForm("username")
	birth := c.PostForm("birth")

	if username == "" || birth == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	c.Redirect(http.StatusFound, "/")
	return
}
