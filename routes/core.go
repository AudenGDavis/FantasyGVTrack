package routes

import (
	"github.com/audengdavis/FantasyGVTrack/controllers"
	"github.com/gin-gonic/gin"
)

func CoreRoutes(r *gin.Engine) {
	coreGroup := r.Group("")
	{
		coreGroup.GET("/", controllers.GetAboutPage)
		coreGroup.GET("/login", controllers.GetLoginPage)
		coreGroup.GET("/signup", controllers.GetSignUpPage)
	}
}
