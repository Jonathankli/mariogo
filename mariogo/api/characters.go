package api

import (
	"jkli/mariogo/mariogo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetCharacters(c *gin.Context) {
	var character []mariogo.Character // Annahme: mariogo.Game ist dein Model für Spiele

	// Lade Spiele mit vorab geladenen Runden und Spielern (nur bestimmte Spalten)
	mariogo.DB.Preload(clause.Associations).Find(&character)

	c.JSON(200, gin.H{
		"games": character,
	})
}

func GetCharacter(c *gin.Context) {
	var character mariogo.Character //siehe models.go
	mariogo.DB.Preload(clause.Associations).First(&character, c.Param("id"))

	if character.ID == 0 {
		c.JSON(404, gin.H{
			"error": "Character not found",
		})
		return
	}

	c.JSON(200, character)
}

func GetPersonsByCharacter(c *gin.Context) {
	var character mariogo.Character

	// Versuchen Sie, den Character anhand der ID zu finden
	if err := mariogo.DB.Preload("Persons").First(&character, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Character not found"})
		return
	}

	c.JSON(200, gin.H{
		"character": character.Name,
		"persons":   character.Persons,
	})
}
