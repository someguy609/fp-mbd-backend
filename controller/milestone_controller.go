package controller

import (
	"fp_mbd/service"

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
