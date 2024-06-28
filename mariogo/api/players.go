package api

import (
	"fmt"
	"jkli/mariogo/mariogo"

	"github.com/gin-gonic/gin"
)

func PatchPlayer(c *gin.Context) {
	var input mariogo.PlayerUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var player mariogo.Player
	gameID := c.Param("game_id")
	numberID := c.Param("number_id")

	if err := mariogo.DB.Where("game_id = ? AND number = ?", gameID, numberID).First(&player).Error; err != nil {
		c.JSON(404, gin.H{"error": "Player not found"})
		return
	}

	if input.FallbackName != nil {
		player.FallbackName = input.FallbackName
	}
	if input.CharacterID != nil {
		player.CharacterID = input.CharacterID
	}
	if input.PersonID != nil {
		player.PersonID = input.PersonID
	}

	if err := mariogo.DB.Save(&player).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update player"})
		return
	}

	c.JSON(200, player)
}
func PatchCurrentPlayer(c *gin.Context) {
	var input mariogo.PlayerUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var player mariogo.Player

	numberID := c.Param("number_id")
	var game mariogo.Game

	// Abrufen der höchsten game_id aus der games Tabelle
	if err := mariogo.DB.Model(&mariogo.Game{}).Where("finished = ?", false).Order("id desc").Limit(1).First(&game).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve game"})
		return
	}
	// Ausgabe der gefundenen game_id (optional)
	fmt.Printf("Gefundene game_id: %d\n", game.ID)

	// Verwenden der abgerufenen game_id in der Datenbankabfrage für den Spieler
	if err := mariogo.DB.Where("game_id = ? AND number = ?", game.ID, numberID).First(&player).Error; err != nil {
		c.JSON(404, gin.H{"error": "Player not found"})
		return
	}

	if input.FallbackName != nil {
		player.FallbackName = input.FallbackName
	}
	if input.CharacterID != nil {
		player.CharacterID = input.CharacterID
	}
	if input.PersonID != nil {
		player.PersonID = input.PersonID
	}

	if err := mariogo.DB.Save(&player).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update player"})
		return
	}

	c.JSON(200, player)
}

func PatchLastPlayer(c *gin.Context) {
	var input mariogo.PlayerUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var player mariogo.Player

	numberID := c.Param("number_id")
	var game mariogo.Game

	// Abrufen der höchsten game_id aus der games Tabelle
	if err := mariogo.DB.Model(&mariogo.Game{}).Order("id desc").Limit(1).First(&game).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve game"})
		return
	}
	// Ausgabe der gefundenen game_id (optional)
	fmt.Printf("Gefundene game_id: %d\n", game.ID)

	// Verwenden der abgerufenen game_id in der Datenbankabfrage für den Spieler
	if err := mariogo.DB.Where("game_id = ? AND number = ?", game.ID, numberID).First(&player).Error; err != nil {
		c.JSON(404, gin.H{"error": "Player not found"})
		return
	}

	if input.FallbackName != nil {
		player.FallbackName = input.FallbackName
	}
	if input.CharacterID != nil {
		player.CharacterID = input.CharacterID
	}
	if input.PersonID != nil {
		player.PersonID = input.PersonID
	}

	if err := mariogo.DB.Save(&player).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update player"})
		return
	}

	c.JSON(200, player)
}
