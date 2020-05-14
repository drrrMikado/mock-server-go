package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	m, err := s.GetMock(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
		"data":    m,
	})
	return
}

func List(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	mlp, err := s.GetMocks(page, pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
		"data":    mlp,
	})
	return
}

func Save(c *gin.Context) {
}

func Update(c *gin.Context) {
}

func Delete(c *gin.Context) {
}
