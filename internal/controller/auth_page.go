package controller

import "github.com/gin-gonic/gin"

func AuthProfile(c *gin.Context) {

	c.HTML(200, "profile.html", gin.H{
		"Message": "WELCOME",
	})

}
