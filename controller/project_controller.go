package controller

import (
	"net/http"
	"strconv"

	"fp_mbd/dto"

	"fp_mbd/service"

	"fp_mbd/utils"

	"github.com/gin-gonic/gin"
)

type (
	ProjectController interface {
		Create(ctx *gin.Context)
		GetProject(ctx *gin.Context)
		GetAllProject(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	projectController struct {
		projectService service.ProjectService
	}
)

func NewProjectController(ps service.ProjectService) ProjectController {
	return &projectController{
		projectService: ps,
	}
}

func (c *projectController) Create(ctx *gin.Context) {
	var project dto.ProjectCreateRequest
	if err := ctx.ShouldBind(&project); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.projectService.Create(ctx.Request.Context(), project)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_PROJECT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_PROJECT, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *projectController) GetAllProject(ctx *gin.Context) {
	var req dto.PaginationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.projectService.GetAllProjectWithPagination(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_PROJECT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_PROJECT,
		Data:    result.Data,
		Meta:    result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *projectController) GetProject(ctx *gin.Context) {
	projectId_, err := strconv.ParseUint(ctx.Param("project_id"), 10, 32)

	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_PROJECT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	projectId := uint(projectId_)

	result, err := c.projectService.GetProjectById(ctx.Request.Context(), projectId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_PROJECT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_PROJECT, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *projectController) Update(ctx *gin.Context) {
	var req dto.ProjectUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	projectId_, err := strconv.ParseUint(ctx.Param("project_id"), 10, 32)

	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_PROJECT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	projectId := uint(projectId_)

	result, err := c.projectService.Update(ctx.Request.Context(), req, projectId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_PROJECT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_PROJECT, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *projectController) Delete(ctx *gin.Context) {
	projectId_, err := strconv.ParseUint(ctx.Param("project_id"), 10, 32)

	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_PROJECT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	projectId := uint(projectId_)

	if err := c.projectService.Delete(ctx.Request.Context(), projectId); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_PROJECT, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_PROJECT, nil)
	ctx.JSON(http.StatusOK, res)
}
