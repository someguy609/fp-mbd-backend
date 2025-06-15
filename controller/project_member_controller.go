package controller

import (
	"fp_mbd/dto"
	"fp_mbd/service"
	"fp_mbd/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	ProjectMemberController interface {
		Create(ctx *gin.Context)
		GetProjectMembers(ctx *gin.Context)
		GetProjectMemberByProjecMemberId(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}
	projectMemberController struct {
		projectMemberService service.ProjectMemberService
	}
)

func NewProjectMemberController(pms service.ProjectMemberService) ProjectMemberController {
	return &projectMemberController{
		projectMemberService: pms,
	}
}

func (c *projectMemberController) Create(ctx *gin.Context) {
	var req dto.ProjectMemberCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	userId := ctx.MustGet("user_id").(string)
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	result, err := c.projectMemberService.Create(ctx.Request.Context(), req, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create project member"})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *projectMemberController) GetProjectMembers(ctx *gin.Context) {
	result, err := c.projectMemberService.GetProjectMembers(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get project members"})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
func (c *projectMemberController) GetProjectMemberByProjecMemberId(ctx *gin.Context) {
	projectMemberId := ctx.Param("projectMemberId")
	if projectMemberId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Project Member ID is required"})
		return
	}
	projectMemberIdUint, err := utils.StringToUint(projectMemberId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Project Member ID"})
		return
	}

	result, err := c.projectMemberService.GetProjectMemberByProjectMemberId(ctx.Request.Context(), projectMemberIdUint)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get project member"})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
func (c *projectMemberController) Update(ctx *gin.Context) {
	var req dto.ProjectMemberUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	projectMemberId := ctx.Param("projectMemberId")
	if projectMemberId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Project Member ID is required"})
		return
	}
	projectMemberIdUint, err := utils.StringToUint(projectMemberId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Project Member ID"})
		return
	}
	userId := ctx.MustGet("user_id").(string)
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	result, err := c.projectMemberService.Update(ctx.Request.Context(), req, projectMemberIdUint, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project member"})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
func (c *projectMemberController) Delete(ctx *gin.Context) {
	projectMemberId := ctx.Param("projectMemberId")
	if projectMemberId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Project Member ID is required"})
		return
	}
	projectMemberIdUint, err := utils.StringToUint(projectMemberId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Project Member ID"})
		return
	}

	if err := c.projectMemberService.Delete(ctx.Request.Context(), projectMemberIdUint); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project member"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Project member deleted successfully"})
}
