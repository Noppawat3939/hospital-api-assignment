package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *gin.Engine {
	r := gin.New()

	r.Use(CORS())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	return r
}
