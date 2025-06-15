package repository

import (
	"context"
	"fp_mbd/entity"

	"gorm.io/gorm"
)

type (
	ProjectMemberRepository interface {
		Create(ctx context.Context, tx *gorm.DB, projectMember entity.ProjectMember) (entity.ProjectMember, error)
		GetProjectMembers(ctx context.Context, tx *gorm.DB, projectId uint) ([]entity.ProjectMember, error)
		GetProjectMemberByProjectMemberId(ctx context.Context, tx *gorm.DB, projectMemberId uint) (entity.ProjectMember, error)
		GetJoinRequests(ctx context.Context, tx *gorm.DB, projectId uint) ([]entity.ProjectMember, error)
		ApproveJoinRequest(ctx context.Context, tx *gorm.DB, projectMemberId uint) (entity.ProjectMember, error)
		Update(ctx context.Context, tx *gorm.DB, projectMember entity.ProjectMember) (entity.ProjectMember, error)
		Delete(ctx context.Context, tx *gorm.DB, projectMemberId uint) error
		IsUserInProject(ctx context.Context, tx *gorm.DB, userId string, projectId uint) (bool, error)
	}
	projectMemberRepository struct {
		db *gorm.DB
	}
)

func NewProjectMemberRepository(db *gorm.DB) ProjectMemberRepository {
	return &projectMemberRepository{
		db: db,
	}
}

func (r *projectMemberRepository) Create(ctx context.Context, tx *gorm.DB, projectMember entity.ProjectMember) (entity.ProjectMember, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&projectMember).Error; err != nil {
		return entity.ProjectMember{}, err
	}

	return projectMember, nil
}
func (r *projectMemberRepository) GetProjectMembers(ctx context.Context, tx *gorm.DB, projectId uint) ([]entity.ProjectMember, error) {
	if tx == nil {
		tx = r.db
	}

	var projectMembers []entity.ProjectMember
	if err := tx.WithContext(ctx).Where("projects_project_id = ? AND is_active = ?", projectId, true).Find(&projectMembers).Error; err != nil {
		return nil, err
	}

	return projectMembers, nil
}
func (r *projectMemberRepository) GetJoinRequests(ctx context.Context, tx *gorm.DB, projectId uint) ([]entity.ProjectMember, error) {
	if tx == nil {
		tx = r.db
	}

	var projectMembers []entity.ProjectMember
	if err := tx.WithContext(ctx).
		Where("projects_project_id = ? AND is_active = ?", projectId, false).
		Find(&projectMembers).Error; err != nil {
		return nil, err
	}

	return projectMembers, nil
}
func (r *projectMemberRepository) ApproveJoinRequest(ctx context.Context, tx *gorm.DB, projectMemberId uint) (entity.ProjectMember, error) {
	if tx == nil {
		tx = r.db
	}

	var projectMember entity.ProjectMember
	if err := tx.WithContext(ctx).Where("project_member_id = ?", projectMemberId).First(&projectMember).Error; err != nil {
		return entity.ProjectMember{}, err
	}

	projectMember.IsActive = true

	if err := tx.WithContext(ctx).Save(&projectMember).Error; err != nil {
		return entity.ProjectMember{}, err
	}

	return projectMember, nil
}

func (r *projectMemberRepository) GetProjectMemberByProjectMemberId(ctx context.Context, tx *gorm.DB, projectMemberId uint) (entity.ProjectMember, error) {
	if tx == nil {
		tx = r.db
	}

	var projectMember entity.ProjectMember
	if err := tx.WithContext(ctx).Where("project_member_id = ?", projectMemberId).Take(&projectMember).Error; err != nil {
		return entity.ProjectMember{}, err
	}

	return projectMember, nil
}

func (r *projectMemberRepository) Update(ctx context.Context, tx *gorm.DB, projectMember entity.ProjectMember) (entity.ProjectMember, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(&projectMember).Error; err != nil {
		return entity.ProjectMember{}, err
	}

	return projectMember, nil
}
func (r *projectMemberRepository) Delete(ctx context.Context, tx *gorm.DB, projectMemberId uint) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Where("project_member_id = ?", projectMemberId).Delete(&entity.ProjectMember{}).Error; err != nil {
		return err
	}

	return nil
}
func (r *projectMemberRepository) IsUserInProject(ctx context.Context, tx *gorm.DB, userId string, projectId uint) (bool, error) {
	// if tx == nil {
	// 	tx = r.db
	// }

	var count int64
	err := r.db.WithContext(ctx).Model(&entity.ProjectMember{}).
		Where("users_user_id = ? AND projects_project_id = ? AND is_active = ?", userId, projectId, true).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
