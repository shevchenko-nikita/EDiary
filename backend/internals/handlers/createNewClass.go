package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/shevchenko-nikita/EDiary/internals/services"
	"net/http"
)

func (h Handler) CreateNewClassHandler(c *gin.Context) {
	teacher, ok := GetUserFromCookie(c)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
	}

	var req struct {
		ClassName string `json:"class_name"`
	}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateNewClass(h.Database, teacher.Id, req.ClassName); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}
