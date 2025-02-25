package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/shevchenko-nikita/EDiary/internals/models"
	"github.com/shevchenko-nikita/EDiary/internals/services"
	"net/http"
	"strconv"
)

func (h Handler) CreateNewClassHandler(c *gin.Context) {
	teacher, ok := GetUserFromCookie(c)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}

	var req struct {
		ClassName string `json:"class_name"`
	}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateNewClass(h.Database, teacher.Id, req.ClassName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

func (h Handler) JoinTheClassHanler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}

	classCode := c.Param("class-code")

	if err := services.JoinTheClass(h.Database, user.Id, classCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error while joining the class"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h Handler) UpdateClassHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}

	var req struct {
		ClassId int    `json:"class_id"`
		NewName string `json:"new_name"`
	}

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateClass(h.Database, user.Id, req.ClassId, req.NewName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error while updating class"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h Handler) DeleteClassHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}

	classId, err := strconv.Atoi(c.Param("class-id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.DeleteClass(h.Database, user.Id, classId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h Handler) LeaveTheClassHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}

	classId, err := strconv.Atoi(c.Param("class-id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.LeaveClass(h.Database, user.Id, classId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
}

func (h Handler) GetUsersListHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}

	classId, err := strconv.Atoi(c.Param("class-id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := services.GetUsersList(h.Database, user.Id, classId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// Get information about teacher from class
func (h Handler) GetClassTeacherHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}

	classId, err := strconv.Atoi(c.Param("class-id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teacher, err := services.GetClassTeacher(h.Database, user.Id, classId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teacher)
}

func (h Handler) CreateAssignmentHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}

	var assignment models.Assignment

	if err := c.ShouldBindWith(&assignment, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateNewAssignment(h.Database, user.Id, &assignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

func (h Handler) DeleteAssignmentHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}

	assignmentId, err := strconv.Atoi(c.Param("assignment-id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.DeleteAssignment(h.Database, user.Id, assignmentId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h Handler) GradeAssignmentHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}

	var mark models.Mark

	if err := c.ShouldBindWith(&mark, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.GradeAssignment(h.Database, user.Id, mark); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
