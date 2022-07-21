package routes

import (
	"github.com/gin-gonic/gin"

	controllers "go-portal/src/controllers"
)

func PublicRoutes(g *gin.RouterGroup) {

	g.GET("/login", controllers.LoginGetHandler())
	g.POST("/login", controllers.LoginPostHandler())

}

func PrivateRoutes(g *gin.RouterGroup) {

	g.GET("/", controllers.IndexGetHandler())
	g.GET("/logout", controllers.LogoutGetHandler())

}
