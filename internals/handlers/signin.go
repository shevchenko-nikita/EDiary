package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/shevchenko-nikita/EDiary/internals/services"
	"net/http"
)

func (h *Handler) SignInHandler(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.SignIn(h.database, req.Username, req.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username or Password is incorrect"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Username and Password are correct"})
}
