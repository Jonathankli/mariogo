package api

import (
	"github.com/gin-gonic/gin"
)

func RunWebServer() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	apiRouter := router.Group("/api")

	apiRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiRouter.GET("/games", GetGames)
	apiRouter.GET("/games/current", GetCurrentGame)
	apiRouter.GET("/games/:id", GetGame)
	apiRouter.PATCH("/games/:game_id/:number_id/player", PatchPlayer)
	apiRouter.PATCH("/games/current/:number_id/player", PatchCurrentPlayer)
	apiRouter.GET("/rounds/:id", GetRound)
	apiRouter.GET("/persons/", GetPersons)
	apiRouter.GET("/persons/:id", GetPerson)
	apiRouter.POST("/persons/", CreatePerson)
	apiRouter.PATCH("/persons/:id", PatchPerson)

	router.NoRoute(func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"test": "test",
		})
	})
	router.Run(":8888")
}
