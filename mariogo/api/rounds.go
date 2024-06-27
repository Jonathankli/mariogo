package api

import (
	"jkli/mariogo/mariogo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetRound(c *gin.Context) {
	var round mariogo.Round //siehe models.go
	mariogo.DB.Preload(clause.Associations).First(&round, c.Param("id"))

	if round.ID == 0 {
		c.JSON(404, gin.H{
			"error": "round not found",
		})
		return
	}

	c.JSON(200, round)
}
