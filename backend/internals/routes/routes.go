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

	router.GET("/profile", middleware.RequireAuth(handler.Database), handler.ProfileHandler)

	classes := router.Group("/classes")
	{
		classes.POST(
			"/create-new-class",
			middleware.RequireAuth(handler.Database),
			handler.CreateNewClassHandler)

		classes.POST(
			"/join-class/:class-code",
			middleware.RequireAuth(handler.Database),
			handler.JoinTheClassHanler)

		classes.PUT(
			"/update-class",
			middleware.RequireAuth(handler.Database),
			handler.UpdateClassHandler)

		classes.DELETE(
			"/delete-class/:class-id",
			middleware.RequireAuth(handler.Database),
			handler.DeleteClassHandler)

		classes.DELETE(
			"/leave-class/:class-id",
			middleware.RequireAuth(handler.Database),
			handler.LeaveTheClassHandler)

		classes.GET(
			"/users-list/:class-id",
			middleware.RequireAuth(handler.Database),
			handler.GetUsersListHandler)

		classes.GET(
			"/teacher/:class-id",
			middleware.RequireAuth(handler.Database),
			handler.GetClassTeacherHandler)

		classes.POST(
			"/create-assignment",
			middleware.RequireAuth(handler.Database),
			handler.CreateAssignmentHandler)

		classes.DELETE(
			"/delete-assignment/:assignment-id",
			middleware.RequireAuth(handler.Database),
			handler.DeleteAssignmentHandler)

		classes.PUT(
			"/grade-assignment",
			middleware.RequireAuth(handler.Database),
			handler.GradeAssignmentHandler)

		classes.GET(
			"/assignments-list/:class-id",
			middleware.RequireAuth(handler.Database),
			handler.GetAssignmentsListHandler)
	}
}
