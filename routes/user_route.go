package routes

import (
	"fp_mbd/constants"
	"fp_mbd/controller"
	"fp_mbd/middleware"
	"fp_mbd/service"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func User(route *gin.Engine, injector *do.Injector) {
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)
	userController := do.MustInvoke[controller.UserController](injector)

	authRoutes := route.Group("/api")
	{
		// Auth
		authRoutes.POST("/register", userController.Register)
		authRoutes.POST("/login", userController.Login)
		authRoutes.GET("/me", middleware.Authenticate(jwtService), userController.Me)
		// authRoutes.POST("/logout", middleware.Authenticate(jwtService), userController.Logout)
	}
	routes := route.Group("/api/user")
	{
		// User
		routes.GET("", middleware.Authenticate(jwtService), userController.GetAllUser) // admin only
		routes.GET("/projects", middleware.Authenticate(jwtService), userController.GetUserProjects)
		routes.GET("/:user_id", userController.GetUserByUserId)                        // admin only
		routes.PATCH("/:user_id", middleware.Authenticate(jwtService), userController.Update)
		routes.DELETE("/:user_id", middleware.Authenticate(jwtService), userController.Delete) // admin only
	}
}
