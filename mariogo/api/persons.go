package api

import (
	"jkli/mariogo/mariogo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetPersons(c *gin.Context) {
	var persons []mariogo.Person // Annahme: mariogo.Game ist dein Model für Spiele

	// Lade Spiele mit vorab geladenen Runden und Spielern (nur bestimmte Spalten)
	mariogo.DB.Preload(clause.Associations).Find(&persons)

	c.JSON(200, gin.H{
		"persons": persons,
	})
}

func GetPerson(c *gin.Context) {
	var person mariogo.Game //siehe models.go
	mariogo.DB.Preload(clause.Associations).First(&person, c.Param("id"))

	if person.ID == 0 {
		c.JSON(404, gin.H{
			"error": "Game not found",
		})
		return
	}

	c.JSON(200, person)
}
