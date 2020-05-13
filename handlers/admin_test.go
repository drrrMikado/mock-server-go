package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type header struct {
	Key   string
	Value string
}

func performRequest(r http.Handler, method, path string, headers ...header) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestGetWithErrorID(t *testing.T) {
	gin.DefaultWriter = os.Stdout
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/test/get/error/:id", func(c *gin.Context) {
		Get(c)
		return
	})
	w := performRequest(router, "GET", "/test/get/error/-1")
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateWithErrorID(t *testing.T) {
	gin.DefaultWriter = os.Stdout
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/test/update/error/:id", func(c *gin.Context) {
		Update(c)
		return
	})
	w := performRequest(router, "POST", "/test/update/error/-1")
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
