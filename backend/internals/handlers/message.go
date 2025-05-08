package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shevchenko-nikita/EDiary/internals/models"
	"github.com/shevchenko-nikita/EDiary/internals/services"
	"net/http"
	"strconv"
)

func (h Handler) CreateClassMessageHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var message models.Message

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message.UserID = user.ID

	if err := services.CreateClassMessage(h.Database, message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h Handler) UpdateMessageHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req struct {
		MessageId int    `json:"message_id"`
		Text      string `json:"text"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateMessage(h.Database, user.ID, req.MessageId, req.Text); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h Handler) DeleteClassMessageHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	messageId, err := strconv.Atoi(c.Param("message-id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.DeleteClassMessage(h.Database, user.ID, messageId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h Handler) GetAllClassMessagesHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	classId, err := strconv.Atoi(c.Param("class-id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	messages, err := services.GetAllClassMessages(h.Database, user.ID, classId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}
