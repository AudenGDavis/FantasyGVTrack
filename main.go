package main

import (
	"strconv"

	"github.com/audengdavis/FantasyGVTrack/models"
	"github.com/audengdavis/FantasyGVTrack/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./static")

	router.LoadHTMLGlob("views/*")

	routes.CoreRoutes(router)

	router.Run(":" + strconv.Itoa(models.Port))
}
