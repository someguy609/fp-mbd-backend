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
		UserID      string `json:"user_id" form:"user_id" binding:"required,min=2,max=15"`
		Name        string `json:"name" form:"name" binding:"required,min=2,max=100"`
		Email       string `json:"email" form:"email" binding:"required,email"`
		Password    string `json:"password" form:"password" binding:"required,min=8"`
		ContactInfo string `json:"contact_info" form:"contact_info" binding:"omitempty,min=8,max=100"`
	}

	CreateMilestoneResponse struct {
		UserID      string `json:"user_id"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		ContactInfo string `json:"contact_info"`
		Role        string `json:"role"`
	}

	GetMilestoneByIdResponse struct {
	}

	UpdateMilestoneRequest struct {
	}

	UpdateMilestoneResponse struct {
		MilestoneID string `json:"milestone_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Status      string `json:"status"`
		StartDate   string `json:"start_date"`
		EndDate     string `json:"end_date"`
	}

	DeleteMilestoneResponse struct {
		MilestoneID string `json:"milestone_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Status      string `json:"status"`
		StartDate   string `json:"start_date"`
		EndDate     string `json:"end_date"`
	}
)
