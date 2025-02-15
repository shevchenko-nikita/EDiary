package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shevchenko-nikita/EDiary/internals/handlers"
)

func InitRoutes(router *gin.Engine, handler *handlers.Handler) {
	router.POST("/sign-up", handler.SignUpHandler)
	router.POST("/sign-in", handler.SignInHandler)
}
