package repository

import (
	"context"
	"fp_mbd/entity"

	"gorm.io/gorm"
)

type (
	ProjectMemberRepository interface {
		Create(ctx context.Context, tx *gorm.DB, projectMember entity.ProjectMember) (entity.ProjectMember, error)
		GetProjectMembersByProjectId(ctx context.Context, tx *gorm.DB, projectId uint) ([]entity.ProjectMember, error)
		GetProjectMemberByUserId(ctx context.Context, tx *gorm.DB, userId string) ([]entity.ProjectMember, error)
		Update(ctx context.Context, tx *gorm.DB, projectMember entity.ProjectMember) (entity.ProjectMember, error)
		Delete(ctx context.Context, tx *gorm.DB, projectMemberId uint) error
		GetProjectMemberById(ctx context.Context, tx *gorm.DB, projectMemberId uint) (entity.ProjectMember, error)
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
func (r *projectMemberRepository) GetProjectMembersByProjectId(ctx context.Context, tx *gorm.DB, projectId uint) ([]entity.ProjectMember, error) {
	if tx == nil {
		tx = r.db
	}

	var projectMembers []entity.ProjectMember
	if err := tx.WithContext(ctx).Where("projects_project_id = ?", projectId).Find(&projectMembers).Error; err != nil {
		return nil, err
	}

	return projectMembers, nil
}
func (r *projectMemberRepository) GetProjectMemberByUserId(ctx context.Context, tx *gorm.DB, userId string) ([]entity.ProjectMember, error) {
	if tx == nil {
		tx = r.db
	}

	var projectMembers []entity.ProjectMember
	if err := tx.WithContext(ctx).Where("users_user_id = ?", userId).Find(&projectMembers).Error; err != nil {
		return nil, err
	}

	return projectMembers, nil
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
func (r *projectMemberRepository) GetProjectMemberById(ctx context.Context, tx *gorm.DB, projectMemberId uint) (entity.ProjectMember, error) {
	if tx == nil {
		tx = r.db
	}

	var projectMember entity.ProjectMember
	if err := tx.WithContext(ctx).Where("project_member_id = ?", projectMemberId).Take(&projectMember).Error; err != nil {
		return entity.ProjectMember{}, err
	}

	return projectMember, nil
}
func (r *projectMemberRepository) IsUserInProject(ctx context.Context, tx *gorm.DB, userId string, projectId uint) (bool, error) {
	if tx == nil {
		tx = r.db
	}

	var count int64
	err := r.db.WithContext(ctx).Model(&entity.ProjectMember{}).
		Where("users_user_id = ? AND projects_project_id = ?", userId, projectId).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
