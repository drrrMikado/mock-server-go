package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	m, err := s.GetMock(id)
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
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	pageSize, _ := strconv.ParseInt(c.Query("pageSize"), 10, 64)
	m, totalPage, err := s.GetMocks(page, pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err_code": 1,
			"err_msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"err_code":   0,
		"err_msg":    "",
		"data":       m,
		"total_page": totalPage,
	})
	return
}

func Save(c *gin.Context) {
}

func Update(c *gin.Context) {
}

func Delete(c *gin.Context) {
}
