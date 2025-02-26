package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shevchenko-nikita/EDiary/internals/models"
	"github.com/shevchenko-nikita/EDiary/internals/services"
	"net/http"
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
