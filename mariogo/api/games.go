package api

import (
	"jkli/mariogo/mariogo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetGames(c *gin.Context) {
	var games []mariogo.Game // Annahme: mariogo.Game ist dein Model für Spiele

	// Lade Spiele mit vorab geladenen Runden und Spielern (nur bestimmte Spalten)
	mariogo.DB.Preload("Players", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, number, game_id, fallback_name") // Hier die gewünschten Spalten angeben
	}).Preload("Placements").Find(&games)

	c.JSON(200, gin.H{
		"games": games,
	})
}

func GetCurrentGame(c *gin.Context) {
	var game mariogo.Game //siehe models.go
	mariogo.DB.Order("id DESC").Preload(clause.Associations).First(&game)

	if game.Finished {
		// Das gefundene Spiel ist bereits beendet
		c.JSON(404, gin.H{"error": "Das Spiel mit der höchsten ID ist bereits beendet"})
		return
	} else {

		c.JSON(200, game)
	}
}

func GetGame(c *gin.Context) {
	var game mariogo.Game //siehe models.go
	mariogo.DB.Preload("Rounds.Placements").Preload(clause.Associations).First(&game, c.Param("id"))

	if game.ID == 0 {
		c.JSON(404, gin.H{
			"error": "Game not found",
		})
		return
	}

	c.JSON(200, game)

	// //Round:
	// var round mariogo.Round //siehe models.go
	// mariogo.DB.Preload(clause.Associations).First(&round, c.Param("id"))

	// if round.ID == 0 {
	// 	c.JSON(404, gin.H{
	// 		"error": "Game not found",
	// 	})
	// 	return
	// }

	// c.JSON(200, round)

}
