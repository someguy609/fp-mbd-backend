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

	routes := route.Group("/api/projects/:project_id")
	{
		routes.POST("/request-join", projectMemberController.Create)
		routes.GET("/members", projectMemberController.GetProjectMembers)
		routes.GET("/join-request", projectMemberController.GetJoinRequests)
		// routes.GET("/:projectMemberId", projectMemberController.GetProjectMemberByProjecMemberId)
		routes.POST("/join-request/:projectMemberId/approve", projectMemberController.ApproveJoinRequest)
		// routes.PATCH("/:projectMemberId", middleware.Authenticate(jwtService), projectMemberController.Update)
		routes.DELETE("/members/:projectMemberId", middleware.Authenticate(jwtService), projectMemberController.Delete)
	}
}
