package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shevchenko-nikita/EDiary/internals/services"
	"net/http"
)

func (h Handler) GetStatisticHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user present"})
		return
	}

	stats, error := services.GetStatisticInfo(h.Database, user.ID)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error})
		return
	}

	c.JSON(http.StatusOK, stats)
}
