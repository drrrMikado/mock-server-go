package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func mockHandler(c *gin.Context) {
	uri := c.Param("uri")
	m, err := svc.GetMockByUriAndMethod(uri, c.Request.Method)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	d, err := time.ParseDuration(m.Delay)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	time.Sleep(d)
	var headers map[string]string
	if err := json.Unmarshal([]byte(m.Headers), &headers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	c.Data(m.StatusCode, headers["Content-Type"], m.Body)
	return
}
