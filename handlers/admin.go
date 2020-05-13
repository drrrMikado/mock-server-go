package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	_ = c.Param("id")
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
