package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/shevchenko-nikita/EDiary/internals/models"
	"github.com/shevchenko-nikita/EDiary/internals/services"
	"net/http"
	"strconv"
)

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

func (h Handler) UpdateAssignmentHandler(c *gin.Context) {
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

	if assignment.TimeCreated != "" || assignment.ClassId != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You can't change class_id and time_created fields"})
		return
	}

	if err := services.UpdateAssignment(h.Database, user.Id, &assignment); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
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

func (h Handler) GetAssignmentsListHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}

	classId, err := strconv.Atoi(c.Param("class-id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong class id"})
		return
	}

	assignments, err := services.GetAssignmentsList(h.Database, user.Id, classId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, assignments)
}

func (h Handler) GetClassTableHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}

	classId, err := strconv.Atoi(c.Param("class-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong class id"})
		return
	}

	assignments, err := services.GetAssignmentsList(h.Database, user.Id, classId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't get assignments list"})
		return
	}

	students, err := services.GetStudentsList(h.Database, user.Id, classId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't get students list"})
		return
	}

	marks, err := services.GetAllClassMarks(h.Database, classId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't get all class_marks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"assignments": assignments,
		"students":    students,
		"marks":       marks,
	})
}
