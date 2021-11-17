package routes

import (
	"github.com/gin-gonic/gin"

	"gin-weight-tracker/app/controllers"
)

func MakeAppRoutes(r *gin.Engine) {
	HomeController := controllers.HomeController{}
	ProfileController := controllers.ProfileController{}
	WeightController := controllers.WeightController{}

	r.GET("/", HomeController.Index)
	r.GET("/login", HomeController.Login)
	r.POST("/login", HomeController.DoLogin)

	r.GET("/profile", ProfileController.Index)
	r.GET("/profile/create", ProfileController.Create)
	r.POST("/profile/store", ProfileController.Store)

	r.GET("/weight/create", WeightController.Create)
	r.POST("/weight/store", WeightController.Store)
}
