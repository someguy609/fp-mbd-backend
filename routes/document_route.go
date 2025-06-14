package routes

import (
	"fp_mbd/constants"
	"fp_mbd/controller"
	"fp_mbd/middleware"
	"fp_mbd/service"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Document(route *gin.Engine, injector *do.Injector) {
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)
	documentController := do.MustInvoke[controller.DocumentController](injector)

	routes := route.Group("/api/document")
	{
		// Document
		// ? auth all routes
		routes.POST("", middleware.Authenticate(jwtService), documentController.Upload)
		routes.GET("", middleware.Authenticate(jwtService), documentController.GetAllDocument)
		routes.DELETE("/:id", middleware.Authenticate(jwtService), documentController.Delete)
		routes.PATCH("", middleware.Authenticate(jwtService), documentController.Update)
		routes.GET("/:id", middleware.Authenticate(jwtService), documentController.GetDocument)
	}
}
