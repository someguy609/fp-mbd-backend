package repository

import (
	"context"
	"fp_mbd/dto"
	"fp_mbd/entity"

	"gorm.io/gorm"
)

type (
	ProjectRepository interface {
		Create(ctx context.Context, tx *gorm.DB, project entity.Project) (entity.Project, error)
		GetAllProjectWithPagination(
			ctx context.Context,
			tx *gorm.DB,
			req dto.PaginationRequest,
		) (dto.GetAllProjectRepositoryResponse, error)
		GetProjectById(ctx context.Context, tx *gorm.DB, projectId uint) (entity.Project, error)
		GetProjectByTitle(ctx context.Context, tx *gorm.DB, title string) (entity.Project, error)
		Update(ctx context.Context, tx *gorm.DB, project entity.Project) (entity.Project, error)
		Delete(ctx context.Context, tx *gorm.DB, projectId uint) error
	}

	projectRepository struct {
		db *gorm.DB
	}
)

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{
		db: db,
	}
}

func (r *projectRepository) Create(ctx context.Context, tx *gorm.DB, project entity.Project) (entity.Project, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&project).Error; err != nil {
		return entity.Project{}, err
	}

	return project, nil
}

func (r *projectRepository) GetAllProjectWithPagination(
	ctx context.Context,
	tx *gorm.DB,
	req dto.PaginationRequest,
) (dto.GetAllProjectRepositoryResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var projects []entity.Project
	var err error
	var count int64

	req.Default()

	query := tx.WithContext(ctx).Model(&entity.Project{})
	if req.Search != "" {
		query = query.Where("title LIKE ?", "%"+req.Search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return dto.GetAllProjectRepositoryResponse{}, err
	}

	if err := query.Scopes(Paginate(req)).Find(&projects).Error; err != nil {
		return dto.GetAllProjectRepositoryResponse{}, err
	}

	totalPage := TotalPage(count, int64(req.PerPage))
	return dto.GetAllProjectRepositoryResponse{
		Projects: projects,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}

func (r *projectRepository) GetProjectById(ctx context.Context, tx *gorm.DB, projectId uint) (entity.Project, error) {
	if tx == nil {
		tx = r.db
	}

	var project entity.Project
	if err := tx.WithContext(ctx).Where("project_id = ?", projectId).Take(&project).Error; err != nil {
		return entity.Project{}, err
	}

	return project, nil
}

func (r *projectRepository) GetProjectByTitle(ctx context.Context, tx *gorm.DB, title string) (entity.Project, error) {
	if tx == nil {
		tx = r.db
	}

	var project entity.Project
	if err := tx.WithContext(ctx).Where("title = ?", title).Take(&project).Error; err != nil {
		return entity.Project{}, err
	}

	return project, nil
}

func (r *projectRepository) Update(ctx context.Context, tx *gorm.DB, project entity.Project) (entity.Project, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&project).Error; err != nil {
		return entity.Project{}, err
	}

	return project, nil
}

func (r *projectRepository) Delete(ctx context.Context, tx *gorm.DB, projectId uint) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.Project{}, "project_id = ?", projectId).Error; err != nil {
		return err
	}

	return nil
}
