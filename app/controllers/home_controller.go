package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type HomeController struct {}

func (h HomeController) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (h HomeController) DoLogin(c *gin.Context) {
	password := c.PostForm("password")
	appPassword := os.Getenv("APP_PASSWORD")
	if password == appPassword {
		plain := []byte(appPassword)
		encryptedHex := md5.Sum(plain)
		encrypted := hex.EncodeToString(encryptedHex[:])

		c.SetCookie("app-token", encrypted, 3600, "", "", false, true)
		c.Redirect(http.StatusFound, "/")
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "Unauthorized",
	})
	return
}

func (h HomeController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
