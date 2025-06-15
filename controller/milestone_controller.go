package controller

import (
	"fp_mbd/dto"
	"fp_mbd/service"
	"fp_mbd/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	MilestoneController interface {
		Create(ctx *gin.Context)
		GetMilestonesByProjectId(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	milestoneController struct {
		milestoneService service.MilestoneService
	}
)

func NewMilestoneController(ms service.MilestoneService) MilestoneController {
	return &milestoneController{
		milestoneService: ms,
	}
}

func (c *milestoneController) Create(ctx *gin.Context) {
	var req dto.MilestoneCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_MILESTONE, err.Error(), nil)
		ctx.AbortWithStatusJSON(400, res)
		return
	}

	user_id := ctx.GetString("user_id")
	if user_id == "" {
		res := utils.BuildResponseFailed(dto.MESSAGE_UNAUTHORIZED_CREATE_MILESTONE, "User ID is required", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	project_id := ctx.Param("project_id")
	if project_id == "" {
		res := utils.BuildResponseFailed("Project ID is required", "Project ID cannot be empty", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	project_id_uint, err := utils.StringToUint(project_id)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid Project ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.milestoneService.Create(ctx.Request.Context(), req, user_id, project_id_uint)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_MILESTONE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_MILESTONE, result)
	ctx.JSON(http.StatusOK, res)

}

func (c *milestoneController) GetMilestonesByProjectId(ctx *gin.Context) {
	projectId := ctx.Param("project_id")
	if projectId == "" {
		res := utils.BuildResponseFailed("Project ID is required", "Project ID cannot be empty", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	projectIdUint, err := utils.StringToUint(projectId)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid Project ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	milestones, err := c.milestoneService.GetMilestonesByProjectId(ctx.Request.Context(), projectIdUint)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_MILESTONE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_MILESTONE, milestones)
	ctx.JSON(http.StatusOK, res)
}

func (c *milestoneController) Update(ctx *gin.Context) {
	var req dto.MilestoneUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed("Failed to get data from body", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	userId := ctx.GetString("user_id")
	if userId == "" {
		res := utils.BuildResponseFailed("User ID is required", "User ID cannot be empty", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	milestoneId := ctx.Param("milestone_id")
	if milestoneId == "" {
		res := utils.BuildResponseFailed("Milestone ID is required", "Milestone ID cannot be empty", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	milestoneIdUint, err := utils.StringToUint(milestoneId)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid Milestone ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.milestoneService.Update(ctx.Request.Context(), req, userId, milestoneIdUint)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_MILESTONE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_MILESTONE, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *milestoneController) Delete(ctx *gin.Context) {
	milestoneId := ctx.Query("milestone_id")
	if milestoneId == "" {
		res := utils.BuildResponseFailed("Milestone ID is required", "Milestone ID cannot be empty", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	milestoneIdUint, err := utils.StringToUint(milestoneId)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid Milestone ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	userId := ctx.MustGet("user_id").(string)

	if err := c.milestoneService.Delete(ctx.Request.Context(), milestoneIdUint, userId); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_MILESTONE, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_MILESTONE, nil)
	ctx.JSON(http.StatusOK, res)
}
