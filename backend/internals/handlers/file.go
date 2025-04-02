package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shevchenko-nikita/EDiary/internals/services"
	"net/http"
	"os"
	"strconv"
)

func (h Handler) UploadFileHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}

	assignment_id, err := strconv.Atoi(c.PostForm("assignment_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid assignment id"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to upload file"})
		return
	}

	fileName := file.Filename
	relativePath, err := SaveFile(c, os.Getenv("FILES_PATH"), file, user.Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to upload file"})
		return
	}

	if err = services.UploadFile(h.Database, fileName, relativePath, user.Id, assignment_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to upload file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
