package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePostHandlerr(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post/"
	r.POST(url, CreatePostHandler)

	body := `{
    "title": "test",
    "content": "for test",
    "community_id": 1
	}`

	req, _ := http.NewRequest(http.MethodGet, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	assert.Contains(t, w.Body.String(), "需要登录")
}
