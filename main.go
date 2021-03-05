package main

import (
	worker "github.com/dubin555/pinger/worker"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	master := worker.NewMaster()
	r.GET("/start/:ip", func(c *gin.Context) {
		ip := c.Param("ip")
		master.Start(ip)
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.GET("/summary/:ip", func(c *gin.Context) {
		ip := c.Param("ip")
		summary := master.Summary(ip)
		c.String(200, summary)
	})
	r.GET("/stop/:ip", func(c *gin.Context) {
		ip := c.Param("ip")
		summary := master.Stop(ip)
		c.String(200, summary)
	})
	r.Run()
}
