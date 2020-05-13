package middlewares

import (
	"regexp"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// method:= c.Request.Method
		origin := c.Request.Header.Get("Origin")
		filterHosts := []string{
			"http://localhost",
			"127.0.0.1",
		}
		for _, v := range filterHosts {
			match, _ := regexp.MatchString(v, origin)
			if match {
				c.Header("Access-Control-Allow-Origin", "*")
				c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
				c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
				c.Set("content-type", "application/json")
			}
		}
		c.Next()
	}
}
