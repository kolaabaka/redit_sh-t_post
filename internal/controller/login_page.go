package controller

import (
	"crypto/rand"
	"goSiteProject/internal/repository"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func LoginFromPage(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")

	userEuserId := repository.CheckUserCreds(login, password)

	if userEuserId == -1 {
		c.HTML(403, "login.html", gin.H{
			"Error": "Login or password is/are wrong",
		})
		return
	}

	bufSecret := make([]byte, 32)

	rand.Read(bufSecret)

	seesionCreated := repository.AddSession(string(bufSecret), userEuserId)

	if seesionCreated {
		c.SetCookie("session_token",
			string(bufSecret[:]),
			3600,
			"/",
			"",
			false,
			false)

		c.Redirect(301, "/")
		return
	}

	c.Redirect(301, "no_page.html")

}

func LogOut(c *gin.Context) {
	var session, _ = c.Cookie("session_token")
	repository.RemoveSession(session)
	c.SetCookie("session_token",
		"",
		-1,
		"/",
		"/",
		false,
		false)
	c.Redirect(301, "/")
}
