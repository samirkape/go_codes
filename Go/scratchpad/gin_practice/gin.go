package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/person/:id", catelog)
	r.Run()
}

func catelog(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		c.URL.path: "User found",
	})
}
