package routes

import (
	"github.com/audengdavis/FantasyGVTrack/controllers"
	"github.com/gin-gonic/gin"
)

func SecurityRoutes(r *gin.Engine) {
	coreGroup := r.Group("/security")
	{
		coreGroup.GET("/test", controllers.GetTest)
	}
}
