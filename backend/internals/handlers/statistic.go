package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) GetStatisticHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user present"})
		return
	}

}
