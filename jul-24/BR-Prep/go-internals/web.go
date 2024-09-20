package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(Middleware())

	r.GET("/get", func(c *gin.Context) {
		c.JSON(http.StatusOK, "All good")
	})
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		st := time.Now()
		c.Next()
		et := time.Now()
		latency := et.Sub(st)
		status := c.Writer.Status()
		method := c.Request.Method

		c.Logger().Infof("%s %s %s", method, status, latency)
	}
}
