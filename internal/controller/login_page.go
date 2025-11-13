package controller

import (
	"crypto/rand"
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func LoginFromPage(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")

	fmt.Println(login + password)

	bufSecret := make([]byte, 32)

	rand.Read(bufSecret)

	c.SetCookie("session_token",
		string(bufSecret[:]),
		3600,
		"/",
		"",
		false,
		false)

	c.Redirect(301, "/")
}
