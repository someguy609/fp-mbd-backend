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
		Create(ctx context.Context, req dto.ProjectMemberCreateRequest, userId string, projectId uint) (dto.ProjectMemberCreateResponse, error)
		GetProjectMembers(ctx context.Context, projectId uint) ([]dto.ProjectMemberGetResponse, error)
		GetJoinRequests(ctx context.Context, projectId uint, userId string) ([]dto.ProjectMemberGetResponse, error)
		ApproveJoinRequest(ctx context.Context, projectMemberId uint, userId string) (dto.ProjectMemberUpdateResponse, error)
		// GetProjectMemberByProjectMemberId(ctx context.Context, projectMemberId uint) (dto.ProjectMemberGetResponse, error)
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

func (s *projectMemberService) Create(ctx context.Context, req dto.ProjectMemberCreateRequest, userId string, projectId uint) (dto.ProjectMemberCreateResponse, error) {
	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return dto.ProjectMemberCreateResponse{}, err
	}
	if user.Role != "mahasiswa" {
		return dto.ProjectMemberCreateResponse{}, dto.ErrUnauthorizedUpdateProjectMember
	}

	projectMember := entity.ProjectMember{
		RoleProject:       req.RoleProject,
		UsersUserID:       userId,
		ProjectsProjectID: projectId,
		IsActive:          false,
	}

	println("Project ID:", projectMember.ProjectsProjectID)

	projectMember, err = s.projectMemberRepo.Create(ctx, nil, projectMember)
	if err != nil {
		return dto.ProjectMemberCreateResponse{}, err
	}

	projectMemberResponse := dto.ProjectMemberCreateResponse{
		ProjectMemberID: projectMember.ProjectMemberID,
		UserID:          projectMember.UsersUserID,
		ProjectID:       projectMember.ProjectsProjectID,
		RoleProject:     projectMember.RoleProject,
		IsActive:        projectMember.IsActive,
		JoinedAt:        projectMember.JoinedAt.String(),
	}

	return projectMemberResponse, nil

}
func (s *projectMemberService) GetProjectMembers(ctx context.Context, projectId uint) ([]dto.ProjectMemberGetResponse, error) {
	projectMembers, err := s.projectMemberRepo.GetProjectMembers(ctx, nil, projectId)
	if err != nil {
		return nil, err
	}

	var projectMemberResponses []dto.ProjectMemberGetResponse
	for _, pm := range projectMembers {
		projectMemberResponses = append(projectMemberResponses, dto.ProjectMemberGetResponse{
			ProjectMemberID: pm.ProjectMemberID,
			UserID:          pm.UsersUserID,
			ProjectID:       pm.ProjectsProjectID,
			RoleProject:     pm.RoleProject,
			IsActive:        pm.IsActive,
			JoinedAt:        pm.JoinedAt.String(),
		})
	}

	return projectMemberResponses, nil
}
func (s *projectMemberService) GetJoinRequests(ctx context.Context, projectId uint, userId string) ([]dto.ProjectMemberGetResponse, error) {
	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return nil, err
	}
	is_inproject, err := s.projectMemberRepo.IsUserInProject(ctx, nil, userId, projectId)
	if err != nil {
		return nil, err
	}

	if user.Role != "dosen" || !is_inproject {
		return nil, dto.ErrUnauthorizedUpdateProjectMember
	}

	projectMembers, err := s.projectMemberRepo.GetJoinRequests(ctx, nil, projectId)
	if err != nil {
		return nil, err
	}
	var projectMemberResponses []dto.ProjectMemberGetResponse
	for _, pm := range projectMembers {
		projectMemberResponses = append(projectMemberResponses, dto.ProjectMemberGetResponse{
			ProjectMemberID: pm.ProjectMemberID,
			UserID:          pm.UsersUserID,
			ProjectID:       pm.ProjectsProjectID,
			RoleProject:     pm.RoleProject,
			IsActive:        pm.IsActive,
			JoinedAt:        pm.JoinedAt.String(),
		})
	}
	return projectMemberResponses, nil
}
func (s *projectMemberService) ApproveJoinRequest(ctx context.Context, projectMemberId uint, userId string) (dto.ProjectMemberUpdateResponse, error) {
	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return dto.ProjectMemberUpdateResponse{}, err
	}
	is_inproject, err := s.projectMemberRepo.IsUserInProject(ctx, nil, userId, projectMemberId)
	if err != nil {
		return dto.ProjectMemberUpdateResponse{}, err
	}
	if user.Role != "dosen" || !is_inproject {
		return dto.ProjectMemberUpdateResponse{}, dto.ErrUnauthorizedUpdateProjectMember
	}
	projectMember, err := s.projectMemberRepo.ApproveJoinRequest(ctx, nil, projectMemberId)
	if err != nil {
		return dto.ProjectMemberUpdateResponse{}, err
	}
	return dto.ProjectMemberUpdateResponse{
		ProjectMemberID: projectMember.ProjectMemberID,
		UserID:          projectMember.UsersUserID,
		ProjectID:       projectMember.ProjectsProjectID,
		RoleProject:     projectMember.RoleProject,
		JoinedAt:        projectMember.JoinedAt.String(),
		IsActive:        projectMember.IsActive,
	}, nil
}

// func (s *projectMemberService) GetProjectMemberByProjectMemberId(ctx context.Context, projectMemberId uint) (dto.ProjectMemberGetResponse, error) {
// 	projectMember, err := s.projectMemberRepo.GetProjectMemberByProjectMemberId(ctx, nil, projectMemberId)
// 	if err != nil {
// 		return dto.ProjectMemberGetResponse{}, err
// 	}

//		return dto.ProjectMemberGetResponse{
//			ProjectMemberID: projectMember.ProjectMemberID,
//			UserID:          projectMember.UserID,
//			ProjectID:       projectMember.ProjectID,
//			RoleProject:     projectMember.RoleProject,
//			JoinedAt:        projectMember.JoinedAt.String(),
//		}, nil
//	}
func (s *projectMemberService) Update(ctx context.Context, req dto.ProjectMemberUpdateRequest, projectMemberId uint, userId string) (dto.ProjectMemberUpdateResponse, error) {

	projectMember, err := s.projectMemberRepo.GetProjectMemberByProjectMemberId(ctx, nil, projectMemberId)
	if err != nil {
		return dto.ProjectMemberUpdateResponse{}, err
	}
	if projectMember.UsersUserID != userId {
		return dto.ProjectMemberUpdateResponse{}, dto.ErrUnauthorizedUpdateProjectMember
	}

	projectMemberEntity := entity.ProjectMember{
		ProjectMemberID:   projectMember.ProjectMemberID,
		RoleProject:       req.RoleProject,
		JoinedAt:          projectMember.JoinedAt,
		UsersUserID:       projectMember.UsersUserID,
		ProjectsProjectID: projectMember.ProjectsProjectID,
	}

	projectMemberRes, err := s.projectMemberRepo.Update(ctx, nil, projectMemberEntity)

	if err != nil {
		return dto.ProjectMemberUpdateResponse{}, err
	}

	return dto.ProjectMemberUpdateResponse{
		ProjectMemberID: projectMemberRes.ProjectMemberID,
		JoinedAt:        projectMemberRes.JoinedAt.String(),
		ProjectID:       projectMemberRes.ProjectsProjectID,
		RoleProject:     projectMemberRes.RoleProject,
		UserID:          projectMemberRes.UsersUserID,
	}, nil
}

func (s *projectMemberService) Delete(ctx context.Context, projectMemberId uint) error {
	err := s.projectMemberRepo.Delete(ctx, nil, projectMemberId)
	if err != nil {
		return err
	}

	return nil
}
