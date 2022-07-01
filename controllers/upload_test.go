package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestUpload(t *testing.T) {
	t.Run("UploadFile header don't have Content-Type correctly", func(t *testing.T) {
		mockResponse := `{"error":"request Content-Type isn't multipart/form-data"}`
		r := SetUpRouter()
		req, _ := http.NewRequest("POST", "/upload", nil)
		r.POST("/upload", UploadFile)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		responseData, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, mockResponse, string(responseData))
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
