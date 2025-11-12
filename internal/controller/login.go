package controller

import (
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
)

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
