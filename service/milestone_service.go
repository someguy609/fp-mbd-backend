package service

import (
	"context"
	"fp_mbd/dto"
	"fp_mbd/repository"

	"gorm.io/gorm"
)

type (
	MilestoneService interface {
		Create(ctx context.Context, req dto.MilestoneCreateRequest) (dto.MilestoneResponse, error)
		GetMilestoneByProjectId(ctx context.Context, projectId string) (dto.MilestonePaginationResponse, error)
		Update(ctx context.Context, req dto.MilestoneUpdateRequest, milestoneId string) (dto.MilestoneResponse, error)
		Delete(ctx context.Context, milestoneId uint) error
	}
	milestoneService struct {
		milestoneRepo repository.MilestoneRepository
		jwtService    JWTService
		db            *gorm.DB
	}
)

func NewMilestoneService(
	milestoneRepo repository.MilestoneRepository,
	jwtService JWTService,
	db *gorm.DB,
) MilestoneService {
	return &milestoneService{
		milestoneRepo: milestoneRepo,
		jwtService:    jwtService,
		db:            db,
	}
}

func (s *milestoneService) Create(ctx context.Context, req dto.MilestoneCreateRequest) (dto.MilestoneResponse, error) {
	milestone, err := s.milestoneRepo.Create(ctx, req)
	if err != nil {
		return dto.MilestoneResponse{}, err
	}

	return dto.ToMilestoneResponse(milestone), nil
}

func (s *milestoneService) GetMilestoneByProjectId(ctx context.Context, projectId string) (dto.MilestonePaginationResponse, error) {
	milestones, err := s.milestoneRepo.GetMilestoneByProjectId(ctx, projectId)
	if err != nil {
		return dto.MilestonePagination
	}
	return dto.ToMilestonePaginationResponse(milestones), nil
}
func (s *milestoneService) Update(ctx context.Context, req dto.MilestoneUpdateRequest, milestoneId string) (dto.MilestoneResponse, error) {
	milestone, err := s.milestoneRepo.Update(ctx, req, milestoneId)
	if err != nil {
		return dto.MilestoneResponse{}, err
	}

	return dto.ToMilestoneResponse(milestone), nil
}
func (s *milestoneService) Delete(ctx context.Context, milestoneId uint) error {
	err := s.milestoneRepo.Delete(ctx, milestoneId)
	if err != nil {
		return err
	}

	return nil
}
