package service

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"fp_mbd/dto"
	"fp_mbd/entity"
	"fp_mbd/repository"
)

type (
	ProjectService interface {
		Create(ctx context.Context, req dto.ProjectCreateRequest, userId string) (dto.ProjectResponse, error)
		GetAllProjectWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.ProjectPaginationResponse, error)
		GetProjectById(ctx context.Context, projectId uint, preload ...string) (dto.ProjectResponse, error)
		GetProjectDocuments(ctx context.Context, projectId uint) (dto.GetProjectDocumentsResponse, error)
		Update(ctx context.Context, req dto.ProjectUpdateRequest, projectId uint) (dto.ProjectUpdateResponse, error)
		Delete(ctx context.Context, projectId uint) error
	}

	projectService struct {
		userRepo          repository.UserRepository
		projectRepo       repository.ProjectRepository
		documentRepo      repository.DocumentRepository
		projectMemberRepo repository.ProjectMemberRepository
		db                *gorm.DB
	}
)

func NewProjectService(
	userRepo repository.UserRepository,
	projectRepo repository.ProjectRepository,
	documentRepo repository.DocumentRepository,
	projectMemberRepo repository.ProjectMemberRepository,
	db *gorm.DB,
) ProjectService {
	return &projectService{
		userRepo:          userRepo,
		projectRepo:       projectRepo,
		documentRepo:      documentRepo,
		projectMemberRepo: projectMemberRepo,
		db:                db,
	}
}

// func SafeRollback(tx *gorm.DB) {
// 	if r := recover(); r != nil {
// 		tx.Rollback()
// 		// TODO: Do you think that we should panic here?
// 		// panic(r)
// 	}
// }

func (s *projectService) Create(ctx context.Context, req dto.ProjectCreateRequest, userId string) (dto.ProjectResponse, error) {

	_, err := s.projectRepo.GetProjectByTitle(ctx, nil, req.Title)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.ProjectResponse{}, err
	}

	project := entity.Project{
		Title:       req.Title,
		Description: req.Description,
		Status:      entity.StatusPlanning,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
	}

	projectReg, err := s.projectRepo.Create(ctx, nil, project)
	if err != nil {
		return dto.ProjectResponse{}, dto.ErrCreateProject
	}

	projectMember := entity.ProjectMember{
		IsActive:          true,
		RoleProject:       dto.ROLE_PROJECT_DOSEN,
		UsersUserID:       userId,
		ProjectsProjectID: projectReg.ProjectID,
	}

	_, err = s.projectMemberRepo.Create(ctx, nil, projectMember)
	if err != nil {
		return dto.ProjectResponse{}, dto.ErrCreateProjectMember
	}

	return dto.ProjectResponse{
		ProjectID:  projectReg.ProjectID,
		Title:      projectReg.Title,
		Status:     projectReg.Status,
		StartDate:  projectReg.StartDate,
		EndDate:    projectReg.EndDate,
		Categories: projectReg.Categories,
	}, nil

}

func (s *projectService) GetAllProjectWithPagination(
	ctx context.Context,
	req dto.PaginationRequest,
) (dto.ProjectPaginationResponse, error) {
	dataWithPaginate, err := s.projectRepo.GetAllProjectWithPagination(ctx, nil, req)
	if err != nil {
		return dto.ProjectPaginationResponse{}, err
	}

	var datas []dto.ProjectResponse
	for _, project := range dataWithPaginate.Projects {
		data := dto.ProjectResponse{
			ProjectID:   project.ProjectID,
			Title:       project.Title,
			Description: project.Description,
			StartDate:   project.StartDate,
			EndDate:     project.EndDate,
			Categories:  project.Categories,
			Status:      project.Status,
		}

		datas = append(datas, data)
	}

	return dto.ProjectPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
func (s *projectService) GetProjectById(ctx context.Context, projectId uint, preload ...string) (dto.ProjectResponse, error) {
	project, err := s.projectRepo.GetProjectById(ctx, nil, projectId, preload...)
	if err != nil {
		return dto.ProjectResponse{}, dto.ErrGetProjectById
	}

	return dto.ProjectResponse{
		ProjectID:  project.ProjectID,
		Title:      project.Title,
		Status:     project.Status,
		StartDate:  project.StartDate,
		EndDate:    project.EndDate,
		Categories: project.Categories,
	}, nil
}

func (s *projectService) GetProjectDocuments(ctx context.Context, projectId uint) (dto.GetProjectDocumentsResponse, error) {
	documents, err := s.documentRepo.GetProjectDocuments(ctx, nil, projectId)
	if err != nil {
		return dto.GetProjectDocumentsResponse{}, err
	}

	return dto.GetProjectDocumentsResponse{
		Documents: documents,
	}, nil
}

func (s *projectService) Update(ctx context.Context, req dto.ProjectUpdateRequest, projectId uint) (
	dto.ProjectUpdateResponse,
	error,
) {
	project, err := s.projectRepo.GetProjectById(ctx, nil, projectId)
	if err != nil {
		return dto.ProjectUpdateResponse{}, dto.ErrProjectNotFound
	}

	data := entity.Project{
		ProjectID:   project.ProjectID,
		StartDate:   project.StartDate,
		EndDate:     project.EndDate,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Categories:  project.Categories,
	}

	projectUpdate, err := s.projectRepo.Update(ctx, nil, data)
	if err != nil {
		return dto.ProjectUpdateResponse{}, dto.ErrUpdateProject
	}

	return dto.ProjectUpdateResponse{
		ProjectID:   projectUpdate.ProjectID,
		Title:       projectUpdate.Title,
		Description: projectUpdate.Description,
		Status:      projectUpdate.Status,
	}, nil
}

func (s *projectService) Delete(ctx context.Context, projectId uint) error {
	tx := s.db.Begin()
	defer SafeRollback(tx)

	project, err := s.projectRepo.GetProjectById(ctx, nil, projectId)
	if err != nil {
		return dto.ErrProjectNotFound
	}

	err = s.projectRepo.Delete(ctx, nil, project.ProjectID)
	if err != nil {
		return dto.ErrDeleteProject
	}

	return nil
}
