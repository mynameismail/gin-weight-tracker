package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type HomeController struct {}

func (h HomeController) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (h HomeController) DoLogin(c *gin.Context) {
	password := c.PostForm("password")
	appPassword := os.Getenv("APP_PASSWORD")
	if password == appPassword {
		key := []byte(appPassword)

		expireIn := 2 * time.Minute

		claims := &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireIn).Unix(),
			Issuer: "app",
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, _ := token.SignedString(key)

		c.SetCookie("app-token", ss, int(expireIn.Seconds()), "", "", false, true)
		c.Redirect(http.StatusFound, "/")
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "Unauthorized",
	})
	return
}

func (h HomeController) Index(c *gin.Context) {
	appPassword := os.Getenv("APP_PASSWORD")
	key := []byte(appPassword)

	cookie, err := c.Cookie("app-token")

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return key, nil
	})
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
	} else {
		fmt.Println(err)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{})
}
