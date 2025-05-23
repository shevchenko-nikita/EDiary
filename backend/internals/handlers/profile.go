package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shevchenko-nikita/EDiary/internals/models"
	"github.com/shevchenko-nikita/EDiary/internals/services"
	"net/http"
	"os"
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

	newUserInfo.ID = user.ID

	if err := services.UpdateUserProfile(h.Database, user.ID, &newUserInfo); err != nil {
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

	if err := services.DeleteProfileImage(h.Database, user.ID); err != nil {
		// TBD
		return
	}

	imgName, err := SaveFile(c, os.Getenv("IMAGE_PATH"), profileImg, user.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't save image"})
		return
	}

	if err := services.UpdateUserProfileImage(h.Database, user.ID, imgName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't save image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h Handler) DeleteProfileImageHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user present"})
		return
	}

	if err := services.DeleteProfileImage(h.Database, user.ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't delete image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
