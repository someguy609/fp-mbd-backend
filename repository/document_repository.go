package repository

import (
	"context"
	"fp_mbd/dto"
	"fp_mbd/entity"

	"gorm.io/gorm"
)

type (
	DocumentRepository interface {
		Upload(ctx context.Context, tx *gorm.DB, document entity.Document) (entity.Document, error)
		GetAllDocumentWithPagination(
			ctx context.Context,
			tx *gorm.DB,
			req dto.PaginationRequest,
		) (dto.GetAllDocumentRepositoryResponse, error)
		GetDocumentById(ctx context.Context, tx *gorm.DB, documentId uint) (entity.Document, error)
		Update(ctx context.Context, tx *gorm.DB, document entity.Document) (entity.Document, error)
		Delete(ctx context.Context, tx *gorm.DB, documentId uint) error
	}

	documentRepository struct {
		db *gorm.DB
	}
)

func NewDocumentRepository(db *gorm.DB) DocumentRepository {
	return &documentRepository{
		db: db,
	}
}

func (r *documentRepository) Upload(ctx context.Context, tx *gorm.DB, document entity.Document) (entity.Document, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&document).Error; err != nil {
		return entity.Document{}, err
	}

	return document, nil
}

func (r *documentRepository) GetAllDocumentWithPagination(
	ctx context.Context,
	tx *gorm.DB,
	req dto.PaginationRequest,
) (dto.GetAllDocumentRepositoryResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var documents []entity.Document
	var err error
	var count int64

	req.Default()

	query := tx.WithContext(ctx).Model(&entity.Document{})
	if req.Search != "" {
		query = query.Where("name LIKE ?", "%"+req.Search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return dto.GetAllDocumentRepositoryResponse{}, err
	}

	if err := query.Scopes(Paginate(req)).Find(&documents).Error; err != nil {
		return dto.GetAllDocumentRepositoryResponse{}, err
	}

	totalPage := TotalPage(count, int64(req.PerPage))
	return dto.GetAllDocumentRepositoryResponse{
		Documents: documents,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}

func (r *documentRepository) GetDocumentById(ctx context.Context, tx *gorm.DB, documentId uint) (entity.Document, error) {
	if tx == nil {
		tx = r.db
	}

	var document entity.Document
	if err := tx.WithContext(ctx).Where("document_id = ?", documentId).Take(&document).Error; err != nil {
		return entity.Document{}, err
	}

	return document, nil
}

func (r *documentRepository) Update(ctx context.Context, tx *gorm.DB, document entity.Document) (entity.Document, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&document).Error; err != nil {
		return entity.Document{}, err
	}

	return document, nil
}

func (r *documentRepository) Delete(ctx context.Context, tx *gorm.DB, documentId uint) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.Document{}, "document_id = ?", documentId).Error; err != nil {
		return err
	}

	return nil
}
