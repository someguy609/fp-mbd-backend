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

	projectRoutes := route.Group("/api/project/:project_id/milestones")
	{
		projectRoutes.POST("", milestoneController.Create)
		projectRoutes.GET("", milestoneController.GetMilestonesByProjectId)
	}
	milestoneRoutes := route.Group("/api/milestones")
	{
		milestoneRoutes.PATCH("/:milestone_id", middleware.Authenticate(jwtService), milestoneController.Update)
		milestoneRoutes.DELETE("/:milestone_id", middleware.Authenticate(jwtService), milestoneController.Delete)
	}
}
