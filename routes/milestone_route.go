package routes

import (
	"fp_mbd/constants"
	"fp_mbd/controller"
	"fp_mbd/middleware"
	"fp_mbd/service"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Milestone(route *gin.Engine, injector *do.Injector) {
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)
	milestoneController := do.MustInvoke[controller.MilestoneController](injector)

	routes := route.Group("/api/milestone")
	{
		routes.POST("", milestoneController.Create)
		routes.GET("/:project_id", milestoneController.GetMilestoneByProjectId)
		routes.PATCH("", middleware.Authenticate(jwtService), milestoneController.Update)
		routes.DELETE("", middleware.Authenticate(jwtService), milestoneController.Delete)
	}
}
