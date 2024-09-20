package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LogggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		st := time.Now()
		c.Next()
		et := time.Now()
		latency := et.Sub(st)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		c.Logger().Infof("%s %s %d %s", method, path, status, latency)
	}
}

func main() {
	r := gin.Default()

	r.Use(LogggerMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}
