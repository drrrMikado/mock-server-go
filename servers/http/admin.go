package http

import (
	"net/http"
	"strconv"

	"github.com/drrrMikado/mock-server-go/models"
	"github.com/gin-gonic/gin"
)

func get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	m, err := svc.GetMock(id)
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

func list(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	mlp, err := svc.GetMocks(page, pageSize)
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

func add(c *gin.Context) {
	mp := &models.AddMockParam{}
	if err := c.ShouldBind(mp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	m, err := svc.AddMock(mp)
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

func update(c *gin.Context) {
	mp := &models.UpdateMockParam{}
	if err := c.ShouldBind(mp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	if err := svc.UpdateMock(mp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
		"data":    gin.H{},
	})
	return
}

func del(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := svc.DeleteMock(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
		"data":    gin.H{},
	})
	return
}
