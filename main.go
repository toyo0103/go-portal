package main

import (
	"go-portal/globals"
	"go-portal/middleware"
	"go-portal/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./view/assets")
	r.LoadHTMLGlob("view/html/*.html")

	r.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := r.Group("/")
	routes.PublicRoutes(public)

	private := r.Group("/")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private)

	// indexRouting := r.Group("/")
	// {
	// 	indexRouting.GET("", handler.GetIndex)
	// }
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
