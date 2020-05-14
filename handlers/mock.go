package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	uri := c.Param("uri")
	m, err := s.GetMockByUriAndMethod(uri, c.Request.Method)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err_code": 1,
			"err_msg":  err.Error(),
		})
		return
	}
	d, err := time.ParseDuration(m.Delay)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err_code": 1,
			"err_msg":  err.Error(),
		})
		return
	}
	time.Sleep(d)
	c.Data(m.StatusCode, m.Headers["Content-Type"], m.Body)
	return
}
