package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/shevchenko-nikita/EDiary/internals/models"
)

type Handler struct {
	Database *sql.DB
}

func NewHandler(database *sql.DB) *Handler {
	return &Handler{Database: database}
}

func GetUserFromCookie(c *gin.Context) (*models.User, bool) {
	userCookie, ok := c.Get("user")

	if !ok {
		return nil, false
	}

	user, ok := userCookie.(*models.User)

	if !ok {
		return nil, false
	}

	return user, true
}
