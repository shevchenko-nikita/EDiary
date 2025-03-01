package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shevchenko-nikita/EDiary/internals/handlers"
	"github.com/shevchenko-nikita/EDiary/internals/middleware"
)

func InitRoutes(router *gin.Engine, handler *handlers.Handler) {
	router.Static("/images", "./images")

	router.POST("/sign-up", handler.SignUpHandler)
	router.POST("/sign-in", handler.SignInHandler)
	router.POST("/logout", handler.LogoutHandler)
	router.PUT(
		"/update-profile",
		middleware.RequireAuth(handler.Database),
		handler.UpdateUserProfileHandler)

	router.GET("/profile", middleware.RequireAuth(handler.Database), handler.ProfileHandler)

	classes := router.Group("/classes")

	classes.Use(middleware.RequireAuth(handler.Database))
	{
		classes.POST("/create-new-class", handler.CreateNewClassHandler)
		classes.POST("/join-class/:class-code", handler.JoinTheClassHanler)
		classes.PUT("/update-class", handler.UpdateClassHandler)
		classes.DELETE("/delete-class/:class-id", handler.DeleteClassHandler)
		classes.DELETE("/leave-class/:class-id", handler.LeaveTheClassHandler)

		classes.GET("/student-list/:class-id", handler.GetStudentsListHandler)
		classes.GET("/teacher/:class-id", handler.GetClassTeacherHandler)

		classes.POST("/create-assignment", handler.CreateAssignmentHandler)
		classes.PUT("/update-assignment", handler.UpdateAssignmentHandler)
		classes.DELETE("/delete-assignment/:assignment-id", handler.DeleteAssignmentHandler)
		classes.GET("/assignments-list/:class-id", handler.GetAssignmentsListHandler)
		classes.PUT("/grade-assignment", handler.GradeAssignmentHandler)

		classes.POST("/create-message", handler.CreateClassMessageHandler)
		classes.PUT("update-message", handler.UpdateMessageHandler)
		classes.DELETE("/delete-message/:message-id", handler.DeleteClassMessageHandler)
		classes.GET("/all-messages/:class-id", handler.GetAllClassMessagesHandler)
	}
}
