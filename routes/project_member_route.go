package routes

import (
	"fp_mbd/constants"
	"fp_mbd/controller"
	"fp_mbd/middleware"
	"fp_mbd/service"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func ProjectMember(route *gin.Engine, injector *do.Injector) {
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)
	projectMemberController := do.MustInvoke[controller.ProjectMemberController](injector)

	routes := route.Group("/api/project_member")
	{
		routes.POST("", projectMemberController.Create)
		routes.GET("", projectMemberController.GetProjectMembers)
		routes.GET("/:projectMemberId", projectMemberController.GetProjectMemberByProjecMemberId)
		routes.PATCH("/:projectMemberId", middleware.Authenticate(jwtService), projectMemberController.Update)
		routes.DELETE("/:projectMemberId", middleware.Authenticate(jwtService), projectMemberController.Delete)
	}
}
