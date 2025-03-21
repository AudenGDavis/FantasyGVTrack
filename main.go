package main

import (
	"github.com/audengdavis/FantasyGVTrack/models"
	"github.com/audengdavis/FantasyGVTrack/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./static")

	router.LoadHTMLGlob("views/*")

	routes.CoreRoutes(router)
	routes.SecurityRoutes(router)

	router.Run(":" + models.Port)

	//database test
	// database.Open("database/main.db")
	// fmt.Println(database.CheckPassword("LegalizeCarrots", "Password123"))
}
