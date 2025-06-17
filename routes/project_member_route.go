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

	routes := route.Group("/api/project/:project_id")
	{
		routes.POST("/request-join", middleware.Authenticate(jwtService), projectMemberController.Create)
		routes.GET("/members", middleware.Authenticate(jwtService), projectMemberController.GetProjectMembers)
		routes.GET("/join-request", middleware.Authenticate(jwtService), projectMemberController.GetJoinRequests)
		// routes.GET("/:projectMemberId", projectMemberController.GetProjectMemberByProjecMemberId)
		routes.PATCH("/join-request/:projectMemberId/approve", middleware.Authenticate(jwtService), projectMemberController.ApproveJoinRequest)
		// routes.PATCH("/:projectMemberId", middleware.Authenticate(jwtService), projectMemberController.Update)
		routes.DELETE("/members/:projectMemberId", middleware.Authenticate(jwtService), projectMemberController.Delete)
	}
}
