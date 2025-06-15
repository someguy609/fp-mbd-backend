package dto

import (
	"errors"
	"mime/multipart"

	"fp_mbd/entity"
)

const (
	// Failed
	MESSAGE_FAILED_UPLOAD_DOCUMENT   = "failed upload documents"
	MESSAGE_FAILED_GET_LIST_DOCUMENT = "failed get list documents"
	MESSAGE_FAILED_GET_DOCUMENT      = "failed get document"
	MESSAGE_FAILED_UPDATE_DOCUMENT   = "failed update document"
	MESSAGE_FAILED_DELETE_DOCUMENT   = "failed delete document"

	// Success
	MESSAGE_SUCCESS_UPLOAD_DOCUMENT   = "success create document"
	MESSAGE_SUCCESS_GET_LIST_DOCUMENT = "success get list document"
	MESSAGE_SUCCESS_GET_DOCUMENT      = "success get document"
	MESSAGE_SUCCESS_DOCUMENT          = "success login"
	MESSAGE_SUCCESS_UPDATE_DOCUMENT   = "success update document"
	MESSAGE_SUCCESS_DELETE_DOCUMENT   = "success delete document"
)

var (
	ErrUploadDocument   = errors.New("failed to upload document")
	ErrGetDocumentById  = errors.New("failed to get document by id")
	ErrUpdateDocument   = errors.New("failed to update document")
	ErrDocumentNotFound = errors.New("document not found")
	ErrDeleteDocument   = errors.New("failed to delete document")
)

type (
	UploadDocumentRequest struct {
		Title     string                `json:"title" form:"title" binding:"required,min=2,max=100"`
		File      *multipart.FileHeader `json:"file" form:"file" binding:"required"`
		ProjectID uint                  `json:"project_id" form:"project_id" binding:"required"`
		// DocumentType    string `json:"document_type" form:"document_type" binding:"required,min=8"` // manual or automated ???
	}

	DocumentResponse struct {
		DocumentID uint   `json:"document_id"`
		Title      string `json:"title"`
		FileURL    string `json:"file_url"`
	}

	DocumentPaginationResponse struct {
		Data []DocumentResponse `json:"data"`
		PaginationResponse
	}

	GetAllDocumentRepositoryResponse struct {
		Documents []entity.Document `json:"documents"`
		PaginationResponse
	}

	DocumentUpdateRequest struct {
		DocumentID uint                  `json:"document_id" form:"document_id"`
		Title      string                `json:"title" form:"title" binding:"omitempty,min=2,max=100"`
		File       *multipart.FileHeader `json:"file" form:"file" binding:"omitempty"`
	}

	DocumentUpdateResponse struct {
		DocumentID uint   `json:"document_id"`
		Title      string `json:"title"`
		FileURL    string `json:"file_url"`
	}
)
