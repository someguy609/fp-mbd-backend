package dto

import (
	"errors"
	"mime/multipart"
	"time"

	"fp_mbd/entity"
)

const (
	// Failed
	MESSAGE_FAILED_CREATE_PROJECT   = "failed create project"
	MESSAGE_FAILED_GET_LIST_PROJECT = "failed get list project"
	MESSAGE_FAILED_GET_PROJECT      = "failed get project"
	MESSAGE_FAILED_UPDATE_PROJECT   = "failed update project"
	MESSAGE_FAILED_DELETE_PROJECT   = "failed delete project"
	// Success
	MESSAGE_SUCCESS_CREATE_PROJECT   = "success create project"
	MESSAGE_SUCCESS_GET_LIST_PROJECT = "success get list project"
	MESSAGE_SUCCESS_GET_PROJECT      = "success get project"
	MESSAGE_SUCCESS_UPDATE_PROJECT   = "success update project"
	MESSAGE_SUCCESS_DELETE_PROJECT   = "success delete project"

	ROLE_PROJECT_DOSEN = "manager"
)

var (
	ErrCreateProject             = errors.New("failed to create project")
	ErrGetProjectById            = errors.New("failed to get project by id")
	ErrUpdateProject             = errors.New("failed to update project")
	ErrProjectNotFound           = errors.New("project not found")
	ErrDeleteProject             = errors.New("failed to delete project")
	ErrUnauthorizedCreateProject = errors.New("unauthorized create project")
)

type (
	ProjectCreateRequest struct {
		Title       string `gorm:"type:varchar(100);not null" json:"title"`
		Description string `gorm:"type:text" json:"description"`
		// Status      string    `gorm:"type:varchar(20);not null" json:"status"`
		StartDate  time.Time `json:"start_date"`
		EndDate    time.Time `json:"end_date"`
		Categories string    `gorm:"type:varchar(100)" json:"categories"` // array of strings ?
	}

	ProjectResponse struct {
		ProjectID   uint                 `json:"project_id"`
		Title       string               `json:"title"`
		Description string               `json:"description"`
		StartDate   time.Time            `json:"start_date"`
		EndDate     time.Time            `json:"end_date"`
		Categories  string               `json:"categories"`
		Status      entity.ProjectStatus `json:"status"`
	}

	ProjectPaginationResponse struct {
		Data []ProjectResponse `json:"data"`
		PaginationResponse
	}

	GetAllProjectRepositoryResponse struct {
		Projects []entity.Project `json:"projects"`
		PaginationResponse
	}

	UploadProjectDocumentRequest struct {
		Title string                `json:"title" form:"title" binding:"required,min=2,max=100"`
		File  *multipart.FileHeader `json:"file" form:"file" binding:"required"`
	}

	GetProjectDocumentRequest struct {
		Data []entity.Document `json:"documents"`
		PaginationResponse
	}

	ProjectUpdateRequest struct {
		Title       string               `gorm:"type:varchar(100);not null" json:"title"`
		Description string               `gorm:"type:text" json:"description"`
		Status      entity.ProjectStatus `gorm:"type:varchar(20);not null" json:"status"`
	}

	ProjectUpdateResponse struct {
		ProjectID   uint                 `json:"project_id"`
		Title       string               `json:"title"`
		Description string               `json:"description"`
		Status      entity.ProjectStatus `json:"status"`
	}
)
