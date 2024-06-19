package mariogo

import "github.com/gin-gonic/gin"

func RunWebServer() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"test": "test",
		})
	})

	router.Run(":8888")
}