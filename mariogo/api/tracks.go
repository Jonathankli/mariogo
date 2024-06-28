// package api

// import (
// 	"jkli/mariogo/mariogo"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm/clause"
// )

// func GetTracks(c *gin.Context) {
// 	var tracks []mariogo.Track // Annahme: mariogo.Game ist dein Model für Spiele

// 	// Lade Spiele mit vorab geladenen Runden und Spielern (nur bestimmte Spalten)
// 	mariogo.DB.Preload(clause.Associations).Find(&tracks)

// 	c.JSON(200, gin.H{
// 		"persons": tracks,
// 	})
// }

// func GetTrack(c *gin.Context) {
// 	var track mariogo.Track //siehe models.go
// 	mariogo.DB.Preload(clause.Associations).First(&track, c.Param("id"))

// 	if track.ID == 0 {
// 		c.JSON(404, gin.H{
// 			"error": "Game not found",
// 		})
// 		return
// 	}

// 	c.JSON(200, track)
// }
