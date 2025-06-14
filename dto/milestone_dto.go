package dto

import (
	"errors"
)

const (
	// Failed
	MESSAGE_FAILED_CREATE_MILESTONE    = "failed create milestone"
	MESSAGEA_FAILED_GET_LIST_MILESTONE = "failed get list milestone"
	MESSAGE_FAILED_UPDATE_MILESTONE    = "failed update milestone"
	MESSAGE_FAILED_DELETE_MILESTONE    = "failed delete milestone"

	// Success
	MESSAGE_SUCCESS_CREATE_MILESTONE   = "success create milestone"
	MESSAGE_SUCCESS_GET_LIST_MILESTONE = "success get list milestone"
	MESSAGE_SUCCESS_UPDATE_MILESTONE   = "success update milestone"
	MESSAGE_SUCCESS_DELETE_MILESTONE   = "success delete milestone"
)

var (
	ErrCreateMilestone  = errors.New("failed to create milestone")
	ErrGetListMilestone = errors.New("failed to milestone id")
	ErrUpdateMilestone  = errors.New("failed to update milestone")
	ErrDeleteMilestone  = errors.New("failed to delete milestone")
)

// belum implement
type (
	MilestoneCreateRequest struct {
		ProjectID   string `json:"project_id" form:"project_id" binding:"required"`
		UserID      string `json:"user_id" form:"user_id" binding:"required,min=2,max=15"`
		Title       string `json:"title" form:"title" binding:"required,min=2,max=100"`
		Description string `json:"description" form:"description" binding:"required,min=2,max=500"`
		DueDate     string `json:"due_date" form:"due_date" binding:"required,datetime=2006-01-02"`
		Status      string `json:"status" form:"status" binding:"required,oneof=in_progress completed"`
	}

	MilestoneCreateResponse struct {
		MilestoneID string `json:"milestone_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		DueDate     string `json:"due_date"`
		Status      string `json:"status"`
	}

	GetMilestoneByIdResponse struct {
		MilestoneID string `json:"milestone_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		DueDate     string `json:"due_date"`
		Status      string `json:"status"`
		CreatedAt   string `json:"created_at"`
	}

	MilestoneUpdateRequest struct {
		MilestoneID string `json:"milestone_id" form:"milestone_id" binding:"required"`
		ProjectID   string `json:"project_id" form:"project_id" binding:"required"`
		Title       string `json:"title" form:"title" binding:"omitempty,min=2,max=100"`
		Description string `json:"description" form:"description" binding:"omitempty,min=2,max=500"`
		DueDate     string `json:"due_date" form:"due_date" binding:"omitempty,datetime=2006-01-02"`
		Status      string `json:"status" form:"status" binding:"omitempty,oneof=in_progress completed"`
	}

	MilestoneUpdateResponse struct {
		MilestoneID string `json:"milestone_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
		DueDate     string `json:"due_date"`
	}

	MilestoneDeleteResponse struct {
		MilestoneID string `json:"milestone_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
		DueDate     string `json:"due_date"`
	}
)
