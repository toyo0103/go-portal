package controllers

import (
	"github.com/gin-contrib/sessions"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-portal/globals"
)

func Login(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	if user != nil {
		c.Redirect(http.StatusMovedPermanently, "/index")
		return
	}
	c.HTML(http.StatusOK, "login.html", gin.H{
		"content": "",
		"user":    user,
	})
}

func LoginPost(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	if user != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Please logout first"})
		return
	}

	account := c.PostForm("account")
	password := c.PostForm("password")

	if account != "toyo" || password != "abc" {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Incorrect account or password"})
		return
	}
	// if helpers.EmptyUserPass(username, password) {
	// 	c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Parameters can't be empty"})
	// 	return
	// }

	session.Set(globals.Userkey, account)
	if err := session.Save(); err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "Failed to save session"})
		return
	}

	c.Redirect(http.StatusFound, "/")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	log.Println("logging out user:", user)

	if user == nil {
		log.Println("Invalid session token")
		return
	}
	session.Delete(globals.Userkey)
	if err := session.Save(); err != nil {
		log.Println("Failed to save session:", err)
		return
	}

	c.Redirect(http.StatusFound, "/login")
}
