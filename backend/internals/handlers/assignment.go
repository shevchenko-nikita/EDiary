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

	if err := services.CreateNewAssignment(h.Database, user.ID, &assignment); err != nil {
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

	//if assignment.TimeCreated != "" || assignment.ClassID != 0 {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "You can't change class_id and time_created fields"})
	//	return
	//}

	if err := services.UpdateAssignment(h.Database, user.ID, &assignment); err != nil {
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

	assignmentID, err := strconv.Atoi(c.Param("assignment-id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad assignment id"})
		return
	}

	if err := services.DeleteAssignment(h.Database, user.ID, assignmentID); err != nil {
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

	if err := services.GradeAssignment(h.Database, user.ID, mark); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h Handler) GetAssignmentHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}

	assignmentID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad assignment id"})
		return
	}

	assignment, err := services.GetAssignment(h.Database, user.ID, assignmentID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't get assignment"})
		return
	}

	c.JSON(http.StatusOK, assignment)
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

	assignments, err := services.GetAssignmentsList(h.Database, user.ID, classId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't get assignments"})
		return
	}

	c.JSON(http.StatusOK, assignments)
}

func (h Handler) GetMarkHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}

	assignmentID, err := strconv.Atoi(c.Param("assignment-id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong assignment id"})
		return
	}

	mark, err := services.GetMark(h.Database, user.ID, assignmentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get mark"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mark": mark})
}

func (h Handler) GetClassTableHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}

	classID, err := strconv.Atoi(c.Param("class-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong class id"})
		return
	}

	assignments, err := services.GetAssignmentsList(h.Database, user.ID, classID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't get assignments list"})
		return
	}

	students, err := services.GetStudentsList(h.Database, user.ID, classID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't get students list"})
		return
	}

	marks, err := services.GetAllClassMarks(h.Database, classID)

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
