package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/shevchenko-nikita/EDiary/internals/models"
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

	token, err := services.SignIn(h.Database, req.Username, req.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username or Password is incorrect"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"Authorization",
		token,
		3600*24*30,
		"",
		"",
		false,
		true)

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) SignUpHandler(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindWith(&user, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.AddNewUser(h.Database, &user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

func (h *Handler) LogoutHandler(c *gin.Context) {
	c.SetCookie(
		"Authorization",
		"",
		-1,
		"",
		"",
		false,
		true)
}
