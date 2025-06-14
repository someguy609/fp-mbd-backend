package dto

import (
	"errors"
)

const (
	// Failed
	MESSAGE_FAILED_CREATE_PROJECT_MEMBER   = "failed create project member"
	MESSAGE_FAILED_GET_PROJECT_MEMBER      = "failed get project member"
	MESSAGE_FAILED_GET_LIST_PROJECT_MEMBER = "failed get list project member"
	MESSAGE_FAILED_UPDATE_PROJECT_MEMBER   = "failed update project member"
	MESSAGE_FAILED_DELETE_PROJECT_MEMBER   = "failed delete project member"

	// Success
	MESSAGE_SUCCESS_CREATE_PROJECT_MEMBER   = "success create project member"
	MESSAGE_SUCCESS_GET_PROJECT_MEMBER      = "success get project member"
	MESSAGE_SUCCESS_GET_LIST_PROJECT_MEMBER = "success get list project member"
	MESSAGE_SUCCESS_UPDATE_PROJECT_MEMBER   = "success update project member"
	MESSAGE_SUCCESS_DELETE_PROJECT_MEMBER   = "success delete project member"
)

var (
	ErrCreateProjectMember             = errors.New("failed to create project member")
	ErrUnauthorizedUpdateProjectMember = errors.New("unauthorized to update project member")
	ErrGetProjectMember                = errors.New("failed to get project member")
	ErrGetListProjectMember            = errors.New("failed to get list project member")
	ErrUpdateProjectMember             = errors.New("failed to update project member")
	ErrDeleteProjectMember             = errors.New("failed to delete project member")
)

// belum implement
type (
	ProjectMemberCreateRequest struct {
		ProjectID   uint   `json:"project_id" form:"project_id" binding:"required"`
		RoleProject string `json:"role_project" form:"role_project" binding:"required"`
	}
	ProjectMemberCreateResponse struct {
		ProjectMemberID uint   `json:"project_member_id"`
		UserID          string `json:"user_id"`
		ProjectID       uint   `json:"project_id"`
		RoleProject     string `json:"role_project"`
		JoinedAt        string `json:"joined_at"`
	}
	ProjectMemberGetResponse struct {
		ProjectMemberID uint   `json:"project_member_id"`
		UserID          string `json:"user_id"`
		ProjectID       uint   `json:"project_id"`
		RoleProject     string `json:"role_project"`
		JoinedAt        string `json:"joined_at"`
	}
	ProjectMemberUpdateRequest struct {
		RoleProject string `json:"role_project" form:"role_project" binding:"required"`
	}
	ProjectMemberUpdateResponse struct {
		ProjectMemberID uint   `json:"project_member_id"`
		UserID          string `json:"user_id"`
		ProjectID       uint   `json:"project_id"`
		RoleProject     string `json:"role_project"`
		JoinedAt        string `json:"joined_at"`
	}
	ProjectMemberDeleteResponse struct {
		ProjectMemberID uint   `json:"project_member_id"`
		UserID          string `json:"user_id"`
		ProjectID       uint   `json:"project_id"`
		RoleProject     string `json:"role_project"`
		JoinedAt        string `json:"joined_at"`
	}
)
