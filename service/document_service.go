package service

import (
	"context"

	"gorm.io/gorm"

	"fp_mbd/dto"
	"fp_mbd/entity"
	"fp_mbd/repository"
)

type (
	DocumentService interface {
		Upload(ctx context.Context, req dto.UploadDocumentRequest, userId string) (dto.DocumentResponse, error)
		GetAllDocumentWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.DocumentPaginationResponse, error)
		GetDocumentById(ctx context.Context, documentId uint) (dto.DocumentResponse, error)
		Update(ctx context.Context, req dto.DocumentUpdateRequest) (dto.DocumentUpdateResponse, error)
		Delete(ctx context.Context, documentId uint) error
	}

	documentService struct {
		documentRepo repository.DocumentRepository
		minioRepo    repository.MinioRepository
		db           *gorm.DB
	}
)

func NewDocumentService(
	documentRepo repository.DocumentRepository,
	minioRepo repository.MinioRepository,
	db *gorm.DB,
) DocumentService {
	return &documentService{
		documentRepo: documentRepo,
		minioRepo:    minioRepo,
		db:           db,
	}
}

// func SafeRollback(tx *gorm.DB) {
// 	if r := recover(); r != nil {
// 		tx.Rollback()
// 		// TODO: Do you think that we should panic here?
// 		// panic(r)
// 	}
// }

func (s *documentService) Upload(ctx context.Context, req dto.UploadDocumentRequest, userId string) (dto.DocumentResponse, error) {
	// todo: use unique identifier for files
	url, err := s.minioRepo.Upload(ctx, req.Title, req.File)

	if err != nil {
		return dto.DocumentResponse{}, dto.ErrUploadDocument
	}

	document := entity.Document{
		Title:             req.Title,
		UsersUserID:       userId,
		ProjectsProjectID: req.ProjectID,
		FileURL:           url,
		DocumentType:      "pdf", // for now
	}

	documentReg, err := s.documentRepo.Upload(ctx, nil, document)
	if err != nil {
		return dto.DocumentResponse{}, dto.ErrUploadDocument
	}

	return dto.DocumentResponse{
		DocumentID: documentReg.DocumentID,
		Title:      documentReg.Title,
		FileURL:    documentReg.FileURL,
	}, nil

}

func (s *documentService) GetAllDocumentWithPagination(
	ctx context.Context,
	req dto.PaginationRequest,
) (dto.DocumentPaginationResponse, error) {
	dataWithPaginate, err := s.documentRepo.GetAllDocumentWithPagination(ctx, nil, req)
	if err != nil {
		return dto.DocumentPaginationResponse{}, err
	}

	var datas []dto.DocumentResponse
	for _, document := range dataWithPaginate.Documents {
		data := dto.DocumentResponse{
			DocumentID: document.DocumentID,
			Title:      document.Title,
			FileURL:    document.FileURL,
		}

		datas = append(datas, data)
	}

	return dto.DocumentPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}

func (s *documentService) GetDocumentById(ctx context.Context, documentId uint) (dto.DocumentResponse, error) {
	document, err := s.documentRepo.GetDocumentById(ctx, nil, documentId)
	if err != nil {
		return dto.DocumentResponse{}, dto.ErrGetDocumentById
	}

	return dto.DocumentResponse{
		DocumentID: document.DocumentID,
		Title:      document.Title,
		FileURL:    document.FileURL,
	}, nil
}

func (s *documentService) Update(ctx context.Context, req dto.DocumentUpdateRequest) (
	dto.DocumentUpdateResponse,
	error,
) {
	document, err := s.documentRepo.GetDocumentById(ctx, nil, req.DocumentID)
	if err != nil {
		return dto.DocumentUpdateResponse{}, dto.ErrDocumentNotFound
	}

	// todo: use unique id for file
	url, err := s.minioRepo.Upload(ctx, req.Title, req.File)

	if err != nil {
		return dto.DocumentUpdateResponse{}, dto.ErrUpdateDocument
	}

	data := entity.Document{
		DocumentID: document.DocumentID,
		Title:      req.Title,
		FileURL:    url,
	}

	documentUpdate, err := s.documentRepo.Update(ctx, nil, data)
	if err != nil {
		return dto.DocumentUpdateResponse{}, dto.ErrUpdateDocument
	}

	return dto.DocumentUpdateResponse{
		DocumentID: documentUpdate.DocumentID,
		Title:      documentUpdate.Title,
		FileURL:    documentUpdate.FileURL,
	}, nil
}

func (s *documentService) Delete(ctx context.Context, documentId uint) error {
	tx := s.db.Begin()
	defer SafeRollback(tx)

	document, err := s.documentRepo.GetDocumentById(ctx, nil, documentId)
	if err != nil {
		return dto.ErrDocumentNotFound
	}

	err = s.documentRepo.Delete(ctx, nil, document.DocumentID)
	if err != nil {
		return dto.ErrDeleteDocument
	}

	return nil
}
