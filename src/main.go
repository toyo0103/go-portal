package main

import (
	"embed"
	"go-portal/src/globals"
	"go-portal/src/middleware"
	"go-portal/src/routes"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//go:embed view/*
var f embed.FS

func main() {
	r := gin.Default()

	fp, _ := fs.Sub(f, "view/assets")
	r.StaticFS("/assets", http.FS(fp))

	templ := template.Must(template.ParseFS(f, "view/html/*.html"))
	r.SetHTMLTemplate(templ)

	r.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := r.Group("/")
	routes.PublicRoutes(public)

	private := r.Group("/")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
