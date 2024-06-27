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

func CreatePerson(c *gin.Context) {
	var input mariogo.Person

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := mariogo.DB.Create(&input).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"person": input,
	})
}

func PatchPerson(c *gin.Context) {
	var person mariogo.Person

	if err := mariogo.DB.First(&person, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Person not found"})
		return
	}

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := mariogo.DB.Save(&person).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"person": person,
	})
}

func DeletePerson(c *gin.Context) {
	var person mariogo.Person

	// Versuchen Sie, die Person anhand der ID zu finden
	if err := mariogo.DB.First(&person, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Person not found"})
		return
	}

	// Soft-Delete der Person
	if err := mariogo.DB.Delete(&person).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Person deleted successfully",
	})
}
