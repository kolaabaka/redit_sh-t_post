package controller

import (
	"goSiteProject/internal/repository"

	"github.com/gin-gonic/gin"
)

func RegistrationFormPage(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")

	repository.AddUser(login, password)

	//Set cookie after registration
	/*bufSecret := make([]byte, 32)

	rand.Read(bufSecret)

	c.SetCookie("session_token",
		string(bufSecret[:]),
		3600,
		"/",
		"",
		false,
		false)*/

	c.Redirect(301, "/")
}
