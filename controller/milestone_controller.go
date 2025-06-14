package controller

import (
	"fp_mbd/dto"
	"fp_mbd/service"
	"fp_mbd/utils"

	"github.com/gin-gonic/gin"
)

type (
	MilestoneController interface {
		Create(ctx *gin.Context)
		GetMilestoneByProjectId(ctx *gin.Context)
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
		res := utils.BuildResponseFailed("Failed to get data from body", err.Error(), nil)
		ctx.AbortWithStatusJSON(400, res)
		return
	}

	user_id := ctx.GetString("user_id")

	result, err := c.milestoneService.Create(ctx.Request.Context(), req, user_id)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to create milestone", err.Error(), nil)
		ctx.JSON(400, res)
		return
	}
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_MILESTONE, result)
	ctx.JSON(200, res)

}

func (c *milestoneController) GetMilestoneByProjectId(ctx *gin.Context) {
	return

}

func (c *milestoneController) Update(ctx *gin.Context) {
	return
}
func (c *milestoneController) Delete(ctx *gin.Context) {
	return
}
