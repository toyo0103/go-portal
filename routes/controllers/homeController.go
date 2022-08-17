package controllers

import (
	"go-portal/globals"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"content": "This is an index page...",
		"user":    user,
	})
}
