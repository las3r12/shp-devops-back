package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"math/rand"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/api/random", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"number": rand.Int(),
		})
	})
	r.Run()
}
