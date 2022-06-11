package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}
	files := form.File["files"]

	// Ensure temp folder exist if not exist create
	if _, err := os.Stat("temp"); err != nil {
		// tempPath := filepath.Join(".", "temp")
		err := os.Mkdir("temp", os.ModeDir)
		if err != nil {
			c.String(http.StatusInternalServerError, "create folder err: %s", err.Error())
			return
		}
	}

	for _, file := range files {
		filename := filepath.Base(file.Filename)
		filepath := filepath.Join("temp", filename)
		if err := c.SaveUploadedFile(file, filepath); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}
	}

	c.String(http.StatusOK, "Uploaded successfully %d files with fields name=%s and email=%s.", len(files), name, email)
}
