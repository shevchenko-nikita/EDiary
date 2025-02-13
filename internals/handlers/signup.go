package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/shevchenko-nikita/EDiary/internals/models"
	"github.com/shevchenko-nikita/EDiary/internals/services"
	"net/http"
)

func (h *Handler) RegisterUserHandler(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindWith(&user, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//c.JSON(http.StatusCreated, user)
	if err := services.AddNewUser(h.database, &user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	}

}
