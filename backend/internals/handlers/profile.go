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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't upload image"})
		return
	}

	rootPath, _ := os.Getwd()
	//imgName := GenerateFileName(filepath.Ext(profileImg.Filename), user.Id)
	imgName := "text.jpg"

	dstFull := filepath.Join(rootPath, os.Getenv("IMAGE_PATH"), imgName)
	dstRelative := os.Getenv("IMAGE_PATH") + filepath.Base(dstFull)

	if err := services.UpdateUserProfileImage(h.Database, user.Id, dstRelative); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't save image"})
		return
	}

	if err := c.SaveUploadedFile(profileImg, dstFull); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't save image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

//func GenerateFileName(extension string, userId int) string {
//
//}
