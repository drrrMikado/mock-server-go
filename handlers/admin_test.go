package handlers

import (
	"bytes"
	"mime/multipart"
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

func TestAddOrUpdateWithEmptyField(t *testing.T) {
	gin.DefaultWriter = os.Stdout
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/test/add/empty/field", func(c *gin.Context) {
		Add(c)
		return
	})
	router.POST("/test/update/empty/field", func(c *gin.Context) {
		Update(c)
		return
	})
	boundary := "--testBoundary"
	body := new(bytes.Buffer)
	mw := multipart.NewWriter(body)
	defer func() {
		_ = mw.Close()
	}()

	assert.NoError(t, mw.SetBoundary(boundary))
	assert.NoError(t, mw.WriteField("description", "description"))
	assert.NoError(t, mw.WriteField("uri", ""))
	assert.NoError(t, mw.WriteField("method", ""))
	assert.NoError(t, mw.WriteField("delay", "20ms"))
	assert.NoError(t, mw.WriteField("status_code", "200"))
	assert.NoError(t, mw.WriteField("headers", ""))
	assert.NoError(t, mw.WriteField("body", "{}"))
	// Test Add
	req, err := http.NewRequest("POST", "/test/add/empty/field", body)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", gin.MIMEMultipartPOSTForm+"; boundary="+boundary)
	w1 := httptest.NewRecorder()
	router.ServeHTTP(w1, req)
	assert.Equal(t, http.StatusBadRequest, w1.Code)
	// Test Update
	assert.NoError(t, mw.WriteField("id", ""))
	req, err = http.NewRequest("POST", "/test/update/empty/field", body)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", gin.MIMEMultipartPOSTForm+"; boundary="+boundary)
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req)
	assert.Equal(t, http.StatusBadRequest, w2.Code)
}
