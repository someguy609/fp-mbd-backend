package service

import (
	"context"
	"errors"
	"fp_mbd/dto"
	"fp_mbd/entity"
	"fp_mbd/repository"
	"time"

	"gorm.io/gorm"
)

type (
	MilestoneService interface {
		Create(ctx context.Context, req dto.MilestoneCreateRequest, userId string) (dto.MilestoneCreateResponse, error)
		GetMilestonesByProjectId(ctx context.Context, projectId uint) ([]dto.GetMilestoneByIdResponse, error)
		Update(ctx context.Context, req dto.MilestoneUpdateRequest, userId string) (dto.MilestoneUpdateResponse, error)
		Delete(ctx context.Context, milestoneId uint, projectId uint, userId string) error
	}
	milestoneService struct {
		milestoneRepo     repository.MilestoneRepository
		userRepo          repository.UserRepository
		projectMemberRepo repository.ProjectMemberRepository
		jwtService        JWTService
		db                *gorm.DB
	}
)

func NewMilestoneService(
	milestoneRepo repository.MilestoneRepository,
	userRepo repository.UserRepository,
	projectMemberRepo repository.ProjectMemberRepository,
	jwtService JWTService,
	db *gorm.DB,
) MilestoneService {
	return &milestoneService{
		milestoneRepo:     milestoneRepo,
		userRepo:          userRepo,
		projectMemberRepo: projectMemberRepo,
		jwtService:        jwtService,
		db:                db,
	}
}

func (s *milestoneService) Create(ctx context.Context, req dto.MilestoneCreateRequest, userId string) (dto.MilestoneCreateResponse, error) {

	user_id, err := s.userRepo.GetUserById(ctx, s.db, userId)

	user_role := user_id.Role
	is_project_member, err := s.projectMemberRepo.IsUserInProject(ctx, s.db, userId, req.ProjectID)

	if user_role != "dosen" || is_project_member == false {
		return dto.MilestoneCreateResponse{}, dto.ErrCreateMilestone
	}

	layout := "2006-01-02"
	parsedDueDate, err := time.Parse(layout, req.DueDate)
	if err != nil {
		return dto.MilestoneCreateResponse{}, errors.New("internal error: failed to parse date after validation")
	}

	milestone := entity.Milestone{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     parsedDueDate,
		Status:      req.Status,
		ProjectID:   req.ProjectID,
	}

	milestone_repo_res, err := s.milestoneRepo.Create(ctx, nil, milestone)

	if err != nil {
		return dto.MilestoneCreateResponse{}, err
	}

	milestone_response := dto.MilestoneCreateResponse{
		MilestoneID: milestone_repo_res.MilestoneID,
		Title:       milestone.Title,
		Description: milestone.Description,
		DueDate:     milestone.DueDate.Format("2006-01-02"),
		Status:      milestone.Status,
	}

	return milestone_response, nil
}

func (s *milestoneService) GetMilestonesByProjectId(ctx context.Context, projectId uint) ([]dto.GetMilestoneByIdResponse, error) {
	milestones, err := s.milestoneRepo.GetMilestonesByProjectId(ctx, nil, projectId)
	if err != nil {
		return nil, err
	}

	var milestoneResponses []dto.GetMilestoneByIdResponse
	for _, milestone := range milestones {
		milestoneResponse := dto.GetMilestoneByIdResponse{
			MilestoneID: milestone.MilestoneID,
			Title:       milestone.Title,
			Description: milestone.Description,
			DueDate:     milestone.DueDate.Format("2006-01-02"),
			Status:      milestone.Status,
			CreatedAt:   milestone.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		milestoneResponses = append(milestoneResponses, milestoneResponse)
	}

	return milestoneResponses, nil
}

func (s *milestoneService) Update(ctx context.Context, req dto.MilestoneUpdateRequest, userId string) (dto.MilestoneUpdateResponse, error) {

	user_id, err := s.userRepo.GetUserById(ctx, s.db, userId)
	if err != nil {
		return dto.MilestoneUpdateResponse{}, err
	}

	user_role := user_id.Role
	is_project_member, err := s.projectMemberRepo.IsUserInProject(ctx, s.db, userId, req.ProjectID)

	if user_role != "admin" || is_project_member == false {
		return dto.MilestoneUpdateResponse{}, dto.ErrUpdateMilestone
	}

	var updatedMilestone entity.Milestone

	if user_role == "mahasiswa" {
		updatedMilestone = entity.Milestone{
			MilestoneID: req.MilestoneID,
			Status:      req.Status,
		}

	} else {
		timeLayout := "2006-01-02"
		parsedDueDate, err := time.Parse(timeLayout, req.DueDate)
		if err != nil {
			return dto.MilestoneUpdateResponse{}, errors.New("internal error: failed to parse date after validation")
		}
		updatedMilestone = entity.Milestone{
			MilestoneID: req.MilestoneID,
			ProjectID:   req.ProjectID,
			Title:       req.Title,
			Description: req.Description,
			DueDate:     parsedDueDate,
			Status:      req.Status,
		}
	}
	milestone, err := s.milestoneRepo.Update(ctx, nil, updatedMilestone)
	if err != nil {
		return dto.MilestoneUpdateResponse{}, dto.ErrUpdateMilestone
	}
	milestoneResponse := dto.MilestoneUpdateResponse{
		MilestoneID: milestone.MilestoneID,
		Title:       milestone.Title,
		Description: milestone.Description,
		Status:      milestone.Status,
		DueDate:     milestone.DueDate.Format("2006-01-02"),
	}
	return milestoneResponse, nil
}
func (s *milestoneService) Delete(ctx context.Context, milestoneId uint, projectId uint, userId string) error {
	user_id, err := s.userRepo.GetUserById(ctx, s.db, userId)
	if err != nil {
		return err
	}

	user_role := user_id.Role
	is_project_member, err := s.projectMemberRepo.IsUserInProject(ctx, s.db, userId, projectId)

	if user_role != "dosen" || is_project_member == false {
		return dto.ErrUpdateMilestone
	}

	err = s.milestoneRepo.Delete(ctx, nil, milestoneId)
	if err != nil {
		return err
	}

	return nil
}
