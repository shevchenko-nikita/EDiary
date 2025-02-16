package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) ProfileHandler(c *gin.Context) {
	user, err := c.Get("user")

	if err {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, user)
}
