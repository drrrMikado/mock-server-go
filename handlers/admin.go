package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	m, err := s.GetMockCfg(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err_code": 1,
			"err_msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "",
		"data":     m,
	})
	return
}

func List(c *gin.Context) {
	_, _ = strconv.ParseInt(c.Query("page"), 10, 64)
	_, _ = strconv.ParseInt(c.Query("pageSize"), 10, 64)
}

func Save(c *gin.Context) {
}

func Update(c *gin.Context) {
}

func Delete(c *gin.Context) {
}
