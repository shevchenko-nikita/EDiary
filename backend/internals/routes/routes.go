package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shevchenko-nikita/EDiary/internals/handlers"
	"github.com/shevchenko-nikita/EDiary/internals/middleware"
	"time"
)

func InitRoutes(router *gin.Engine, handler *handlers.Handler) {
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:4200"}, // Р
		// азрешенные домены
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Static("/images", "./uploads/images")

	router.POST("/sign-up", handler.SignUpHandler)
	router.POST("/sign-in", handler.SignInHandler)
	router.POST("/logout", handler.LogoutHandler)
	router.GET("/check-auth", middleware.RequireAuth(handler.Database), handler.CheckAuthHandler)

	user := router.Group("/user")

	user.Use(middleware.RequireAuth(handler.Database))
	{
		user.PUT("/update-profile", handler.UpdateUserProfileHandler)
		user.PUT("/update-profile-image", handler.UpdateProfileImageHandler)
		user.DELETE("/delete-profile-image", handler.DeleteProfileImageHandler)
		user.GET("/profile", handler.ProfileHandler)
	}

	classes := router.Group("/classes")

	classes.Use(middleware.RequireAuth(handler.Database))
	{
		classes.POST("/create-new-class", handler.CreateNewClassHandler)
		classes.POST("/join-class/:class-code", handler.JoinTheClassHanler)
		classes.PUT("/update-class", handler.UpdateClassHandler)
		classes.DELETE("/delete-class/:class-id", handler.DeleteClassHandler)
		classes.DELETE("/leave-class/:class-id", handler.LeaveTheClassHandler)
		classes.GET("/education-list", handler.GetEducationClassesHandler)
		classes.GET("teaching-list", handler.GetTeachingListHandler)
		classes.GET("get-info/:class-id", handler.GetClassInfoHandler)

		classes.GET("/student-list/:class-id", handler.GetStudentsListHandler)
		classes.GET("/teacher/:class-id", handler.GetClassTeacherHandler)

		classes.POST("/create-assignment", handler.CreateAssignmentHandler)
		classes.PUT("/update-assignment", handler.UpdateAssignmentHandler)
		classes.DELETE("/delete-assignment/:assignment-id", handler.DeleteAssignmentHandler)
		classes.GET("/assignments-list/:class-id", handler.GetAssignmentsListHandler)
		classes.PUT("/grade-assignment", handler.GradeAssignmentHandler)

		//classes.POST("/upload-file", handler.UploadFileHandler)

		classes.GET("/table/:class-id", handler.GetClassTableHandler)

		classes.POST("/create-message", handler.CreateClassMessageHandler)
		classes.PUT("update-message", handler.UpdateMessageHandler)
		classes.DELETE("/delete-message/:message-id", handler.DeleteClassMessageHandler)
		classes.GET("/all-messages/:class-id", handler.GetAllClassMessagesHandler)
	}
}
