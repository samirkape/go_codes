package main

import (
	"github.com/gin-gonic/gin"
)

// /packages
//	GET/POST
// /packages/id
//  GET

func main() {
	r := gin.Default()
	r.GET("/packages", packagesHandler)
	r.GET("/packages/:id", packageIdHandler)
	r.GET("/packages/:id", packageIdHandler)
	r.Run()
}

func packagesHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Packeges called"})
}

func packageIdHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":     "Packeges id called",
		"id":          c.Param("id"),
		"all_queries": c.QueryArray("make"),
	})
}
