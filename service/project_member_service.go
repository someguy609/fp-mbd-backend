package service

import (
	"context"
	"fp_mbd/dto"
	"fp_mbd/entity"
	"fp_mbd/repository"

	"gorm.io/gorm"
)

type (
	ProjectMemberService interface {
		Create(ctx context.Context, req dto.ProjectMemberCreateRequest, userId string) (dto.ProjectMemberCreateResponse, error)
		GetProjectMembers(ctx context.Context) ([]dto.ProjectMemberGetResponse, error)
		GetProjectMemberByProjectMemberId(ctx context.Context, projectMemberId uint) (dto.ProjectMemberGetResponse, error)
		Update(ctx context.Context, req dto.ProjectMemberUpdateRequest, projectMemberId uint, userId string) (dto.ProjectMemberUpdateResponse, error)
		Delete(ctx context.Context, projectMemberId uint) error
	}

	projectMemberService struct {
		userRepo          repository.UserRepository
		projectMemberRepo repository.ProjectMemberRepository
		jwtService        JWTService
		db                *gorm.DB
	}
)

func NewProjectMemberService(
	userRepo repository.UserRepository,
	projectMemberRepo repository.ProjectMemberRepository,
	jwtService JWTService,
	db *gorm.DB,
) ProjectMemberService {
	return &projectMemberService{
		userRepo:          userRepo,
		projectMemberRepo: projectMemberRepo,
		jwtService:        jwtService,
		db:                db,
	}
}

func (s *projectMemberService) Create(ctx context.Context, req dto.ProjectMemberCreateRequest, userId string) (dto.ProjectMemberCreateResponse, error) {

	projectMember := entity.ProjectMember{
		RoleProject: req.RoleProject,
		ProjectID:   req.ProjectID,
		UserID:      userId,
	}

	projectMember, err := s.projectMemberRepo.Create(ctx, nil, projectMember)
	if err != nil {
		return dto.ProjectMemberCreateResponse{}, err
	}

	projectMemberResponse := dto.ProjectMemberCreateResponse{
		ProjectMemberID: projectMember.ProjectMemberID,
		UserID:          projectMember.UserID,
		ProjectID:       projectMember.ProjectID,
		RoleProject:     projectMember.RoleProject,
		JoinedAt:        projectMember.JoinedAt.String(),
	}

	return projectMemberResponse, nil

}
func (s *projectMemberService) GetProjectMembers(ctx context.Context) ([]dto.ProjectMemberGetResponse, error) {
	projectMembers, err := s.projectMemberRepo.GetProjectMembers(ctx, nil)
	if err != nil {
		return nil, err
	}

	var projectMemberResponses []dto.ProjectMemberGetResponse
	for _, pm := range projectMembers {
		projectMemberResponses = append(projectMemberResponses, dto.ProjectMemberGetResponse{
			ProjectMemberID: pm.ProjectMemberID,
			UserID:          pm.UserID,
			ProjectID:       pm.ProjectID,
			RoleProject:     pm.RoleProject,
			JoinedAt:        pm.JoinedAt.String(),
		})
	}

	return projectMemberResponses, nil
}
func (s *projectMemberService) GetProjectMemberByProjectMemberId(ctx context.Context, projectMemberId uint) (dto.ProjectMemberGetResponse, error) {
	projectMember, err := s.projectMemberRepo.GetProjectMemberByProjectMemberId(ctx, nil, projectMemberId)
	if err != nil {
		return dto.ProjectMemberGetResponse{}, err
	}

	return dto.ProjectMemberGetResponse{
		ProjectMemberID: projectMember.ProjectMemberID,
		UserID:          projectMember.UserID,
		ProjectID:       projectMember.ProjectID,
		RoleProject:     projectMember.RoleProject,
		JoinedAt:        projectMember.JoinedAt.String(),
	}, nil
}
func (s *projectMemberService) Update(ctx context.Context, req dto.ProjectMemberUpdateRequest, projectMemberId uint, userId string) (dto.ProjectMemberUpdateResponse, error) {

	projectMember, err := s.projectMemberRepo.GetProjectMemberByProjectMemberId(ctx, nil, projectMemberId)
	if err != nil {
		return dto.ProjectMemberUpdateResponse{}, err
	}
	if projectMember.UserID != userId {
		return dto.ProjectMemberUpdateResponse{}, dto.ErrUnauthorizedUpdateProjectMember
	}

	projectMemberEntity := entity.ProjectMember{
		ProjectMemberID: projectMember.ProjectMemberID,
		RoleProject:     req.RoleProject,
		JoinedAt:        projectMember.JoinedAt,
		UserID:          projectMember.UserID,
		ProjectID:       projectMember.ProjectID,
	}

	projectMemberRes, err := s.projectMemberRepo.Update(ctx, nil, projectMemberEntity)

	if err != nil {
		return dto.ProjectMemberUpdateResponse{}, err
	}

	return dto.ProjectMemberUpdateResponse{
		ProjectMemberID: projectMemberRes.ProjectMemberID,
		JoinedAt:        projectMemberRes.JoinedAt.String(),
		ProjectID:       projectMemberRes.ProjectID,
		RoleProject:     projectMemberRes.RoleProject,
		UserID:          projectMemberRes.UserID,
	}, nil
}

func (s *projectMemberService) Delete(ctx context.Context, projectMemberId uint) error {
	err := s.projectMemberRepo.Delete(ctx, nil, projectMemberId)
	if err != nil {
		return err
	}

	return nil
}
