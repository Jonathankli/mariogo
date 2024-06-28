package api

import (
	"jkli/mariogo/mariogo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetGames(c *gin.Context) {
	var games []mariogo.Game // Annahme: mariogo.Game ist dein Model für Spiele

	// Lade Spiele mit vorab geladenen Runden und Spielern (nur bestimmte Spalten)
	mariogo.DB.Preload("Rounds", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Placements", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, round_id, player_id, position")
		}).Preload("PlacementChangeLogs", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, time, player1, player2, player3, player4")
		}).Select("id, `index`, track_name, game_id")
	}).Preload("Players", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, number, game_id, fallback_name")
	}).Find(&games)

	c.JSON(200, gin.H{
		"games": games,
	})
}

func GetCurrentGame(c *gin.Context) {
	var game mariogo.Game // siehe models.go
	mariogo.DB.Preload("Rounds", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Placements", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, round_id, player_id, position")
		}).Preload("PlacementChangeLogs", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, round_id, time, player1, player2, player3, player4")
		}).Select("id, `index`, track_name, game_id")
	}).Preload("Players", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, number, game_id, fallback_name")
	}).Where("finished = ?", false).Order("id desc").First(&game)

	if game.ID == 0 {
		c.JSON(404, gin.H{
			"error": "game not found - seems like all games are finished",
		})
		return
	}

	c.JSON(200, game)
}

func GetLastGame(c *gin.Context) {
	var game mariogo.Game // siehe models.go
	mariogo.DB.Preload("Rounds", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Placements", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, round_id, player_id, position")
		}).Preload("PlacementChangeLogs", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, round_id, time, player1, player2, player3, player4")
		}).Select("id, `index`, track_name, game_id")
	}).Preload("Players", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, number, game_id, fallback_name")
	}).Order("id desc").First(&game)

	if game.ID == 0 {
		c.JSON(404, gin.H{
			"error": "game not found - seems like all games are finished",
		})
		return
	}

	c.JSON(200, game)
}

func GetGame(c *gin.Context) {
	var game mariogo.Game //siehe models.go
	mariogo.DB.Preload("Rounds", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Placements", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, round_id, player_id, position")
		}).Preload("PlacementChangeLogs", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, round_id, time, player1, player2, player3, player4")
		}).Select("id, `index`, track_name, game_id")
	}).Preload("Players", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, number, game_id, fallback_name")
	}).First(&game, c.Param("id"))

	if game.ID == 0 {
		c.JSON(404, gin.H{
			"error": "Game not found",
		})
		return
	}

	c.JSON(200, game)
}
