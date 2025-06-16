package routes

import (
	"fp_mbd/constants"
	"fp_mbd/controller"
	"fp_mbd/middleware"
	"fp_mbd/service"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Project(route *gin.Engine, injector *do.Injector) {
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)
	projectController := do.MustInvoke[controller.ProjectController](injector)

	routes := route.Group("/api/project")
	{
		// Project
		// todo: convert roles into enum
		routes.POST("", middleware.Authenticate(jwtService), middleware.RequireRoles("dosen"), projectController.Create)
		routes.GET("", middleware.Authenticate(jwtService), projectController.GetAllProject)
		routes.GET("/:id", middleware.Authenticate(jwtService), projectController.GetProject)
		routes.PATCH("/:id", middleware.Authenticate(jwtService), middleware.RequireRoles("dosen"), projectController.Update)
		routes.DELETE("/:id", middleware.Authenticate(jwtService), middleware.RequireRoles("dosen"), projectController.Delete)
		routes.POST("/:id/documents", middleware.Authenticate(jwtService), projectController.UploadDocument)
		routes.GET("/:id/documents", middleware.Authenticate(jwtService), projectController.GetDocument)
	}
}
