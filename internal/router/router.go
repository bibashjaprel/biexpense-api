package router

import "github.com/gin-gonic/gin"

func Setup() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "biexpense-api is running",
		})
	})

	api := r.Group("/api/v1")

	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})
	}

	return r
}
