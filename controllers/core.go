package controllers

import (
	"net/http"

	"github.com/audengdavis/FantasyGVTrack/models"
	"github.com/gin-gonic/gin"
)

func GetErrorPage(c *gin.Context) {
	c.JSON(501, gin.H{
		"error": "endpoint not implemented",
	})
}

func GetAboutPage(c *gin.Context) {
	println("requested about")
	c.HTML(http.StatusOK, "main.html", gin.H{
		"title":     "Home Page",
		"idAddress": models.IpAddress,
		"port":      models.Port,
	})
}

func GetSignUpPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signUp.html", gin.H{
		"title":     "Home Page",
		"idAddress": models.IpAddress,
		"port":      models.Port,
	})
}

func GetLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title":     "Login",
		"idAddress": models.IpAddress,
		"port":      models.Port,
	})
}
