package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

	if err := services.CreateNewClass(h.Database, teacher.ID, req.ClassName); err != nil {
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

	if err := services.JoinTheClass(h.Database, user.ID, classCode); err != nil {
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

	if err := services.UpdateClass(h.Database, user.ID, req.ClassId, req.NewName); err != nil {
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

	classID, err := strconv.Atoi(c.Param("class-id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.DeleteClass(h.Database, user.ID, classID); err != nil {
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

	classID, err := strconv.Atoi(c.Param("class-id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.LeaveClass(h.Database, user.ID, classID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
}

func (h Handler) GetStudentsListHandler(c *gin.Context) {
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

	students, err := services.GetStudentsList(h.Database, user.ID, classId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

// Get information about teacher from class
func (h Handler) GetClassTeacherHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}

	classID, err := strconv.Atoi(c.Param("class-id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teacher, err := services.GetClassTeacher(h.Database, user.ID, classID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teacher)
}

func (h Handler) GetEducationClassesHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}

	classes, err := services.GetEducationClasses(h.Database, user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(http.StatusOK, classes)
}

func (h Handler) GetTeachingListHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}

	classes, err := services.GetTeachingClasses(h.Database, user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(http.StatusOK, classes)
}

func (h Handler) GetClassInfoHandler(c *gin.Context) {
	classID, err := strconv.Atoi(c.Param("class-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't get class ID"})
		return
	}

	class, err := services.GetClassInfo(h.Database, classID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(http.StatusOK, class)
}

func (h Handler) IsTeacherHandler(c *gin.Context) {
	user, ok := GetUserFromCookie(c)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}

	classID, err := strconv.Atoi(c.Param("class-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't get class ID"})
		return
	}

	class, err := services.GetClassInfo(h.Database, classID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	isTeacher := class.TeacherId == user.ID
	c.JSON(http.StatusOK, gin.H{"isTeacher": isTeacher})
}
