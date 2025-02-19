package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shevchenko-nikita/EDiary/internals/handlers"
	"github.com/shevchenko-nikita/EDiary/internals/middleware"
)

func InitRoutes(router *gin.Engine, handler *handlers.Handler) {
	router.POST("/sign-up", handler.SignUpHandler)
	router.POST("/sign-in", handler.SignInHandler)

	router.GET("/profile", middleware.RequireAuth(handler.Database), handler.ProfileHandler)

	classes := router.Group("/classes")
	{
		classes.POST(
			"/create-new-class",
			middleware.RequireAuth(handler.Database),
			handler.CreateNewClassHandler)

		classes.POST(
			"/join-class",
			middleware.RequireAuth(handler.Database),
			handler.JoinTheClassHanler)
	}
}
