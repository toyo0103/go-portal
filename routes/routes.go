package routes

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"go-portal/globals"
	"go-portal/middleware"
	controllers "go-portal/routes/controllers"
)

var F embed.FS

func InitRouter() *gin.Engine {
	r := gin.Default()
	fp, _ := fs.Sub(F, "view/assets")
	r.StaticFS("/assets", http.FS(fp))

	templ := template.Must(template.ParseFS(F, "view/html/*.html"))
	r.SetHTMLTemplate(templ)

	r.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := r.Group("/")
	publicRoutes(public)

	private := r.Group("/")
	private.Use(middleware.AuthRequired)
	privateRoutes(private)
	return r
}

func publicRoutes(g *gin.RouterGroup) {
	g.GET("/login", controllers.Login)
	g.POST("/login", controllers.LoginPost)
}

func privateRoutes(g *gin.RouterGroup) {
	g.GET("/", controllers.Index)
	g.GET("/logout", controllers.Logout)
}
