package repository

import (
	"context"
	"fp_mbd/entity"

	"gorm.io/gorm"
)

type (
	MilestoneRepository interface {
		Create(ctx context.Context, tx *gorm.DB, milestone entity.Milestone) (entity.Milestone, error)
		GetProjectIdByMilestoneId(ctx context.Context, tx *gorm.DB, milestoneId uint) (uint, error)
		GetMilestonesByProjectId(ctx context.Context, tx *gorm.DB, projectId uint) ([]entity.Milestone, error)
		Update(ctx context.Context, tx *gorm.DB, milestone entity.Milestone) (entity.Milestone, error)
		Delete(ctx context.Context, tx *gorm.DB, milestoneId uint) error
	}

	milestoneRepository struct {
		db *gorm.DB
	}
)

func NewMilestoneRepository(db *gorm.DB) MilestoneRepository {
	return &milestoneRepository{
		db: db,
	}
}

func (r *milestoneRepository) Create(ctx context.Context, tx *gorm.DB, milestone entity.Milestone) (entity.Milestone, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&milestone).Error; err != nil {
		return entity.Milestone{}, err
	}

	return milestone, nil
}
func (r *milestoneRepository) GetMilestonesByProjectId(ctx context.Context, tx *gorm.DB, projectId uint) ([]entity.Milestone, error) {
	if tx == nil {
		tx = r.db
	}

	var milestones []entity.Milestone
	if err := tx.WithContext(ctx).Where("projects_project_id = ?", projectId).Find(&milestones).Error; err != nil {
		println("Error retrieving milestones:", err)
		return nil, err
	}

	println("Retrieved milestones:", len(milestones))

	return milestones, nil
}
func (r *milestoneRepository) Update(ctx context.Context, tx *gorm.DB, milestone entity.Milestone) (entity.Milestone, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(&milestone).Error; err != nil {
		return entity.Milestone{}, err
	}

	return milestone, nil
}
func (r *milestoneRepository) Delete(ctx context.Context, tx *gorm.DB, milestoneId uint) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Where("milestone_id = ?", milestoneId).Delete(&entity.Milestone{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *milestoneRepository) GetProjectIdByMilestoneId(ctx context.Context, tx *gorm.DB, milestoneId uint) (uint, error) {
	if tx == nil {
		tx = r.db
	}

	var milestone entity.Milestone
	if err := tx.WithContext(ctx).Select("projects_project_id").Where("milestone_id = ?", milestoneId).First(&milestone).Error; err != nil {
		return 0, err
	}

	return milestone.ProjectsProjectID, nil
}
