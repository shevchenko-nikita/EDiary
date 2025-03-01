package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shevchenko-nikita/EDiary/internals/models"
	"github.com/shevchenko-nikita/EDiary/internals/services"
	"net/http"
	"os"
	"path/filepath"
)

func (h Handler) ProfileHandler(c *gin.Context) {
	user, ok := c.Get("user")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user present"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h Handler) UpdateUserProfileHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user present"})
		return
	}

	var newUserInfo models.User

	if err := c.BindJSON(&newUserInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateUserProfile(h.Database, user.Id, &newUserInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h Handler) UpdateProfileImageHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user present"})
		return
	}

	profileImg, err := c.FormFile("profile_image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось загрузить файл"})
		return
	}

	rootPath, _ := os.Getwd()
	dstFull := filepath.Join(rootPath, os.Getenv("IMAGE_PATH"), "new_img.jpg")
	dstRelative := os.Getenv("IMAGE_PATH") + filepath.Base(dstFull)

	if err := services.UpdateUserProfileImage(h.Database, user.Id, dstRelative); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't save image"})
		return
	}

	if err := c.SaveUploadedFile(profileImg, dstFull); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
