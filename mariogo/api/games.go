package api

import (
	"jkli/mariogo/mariogo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetGames(c *gin.Context) {

	var games []mariogo.Game
	mariogo.DB.Find(&games)

	c.JSON(200, gin.H{
		"games": games,
	})
}

func GetCurrentGame(c *gin.Context) {
	var game mariogo.Game
	mariogo.DB.Where("finished = ?", false).Preload(clause.Associations).Last(&game)

	c.JSON(200, game)
}

func GetGame(c *gin.Context) {
	var game mariogo.Game
	mariogo.DB.Preload(clause.Associations).First(&game, c.Param("id"))

	if game.ID == 0 {
		c.JSON(404, gin.H{
			"error": "Game not found",
		})
		return
	}

	c.JSON(200, game)
}
