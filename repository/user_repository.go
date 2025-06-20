package repository

import (
	"context"
	"fp_mbd/dto"
	"fp_mbd/entity"
	"log"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Register(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error)
		GetAllUserWithPagination(
			ctx context.Context,
			tx *gorm.DB,
			req dto.PaginationRequest,
		) (dto.GetAllUserRepositoryResponse, error)
		GetUserById(ctx context.Context, tx *gorm.DB, userId string) (entity.User, error)
		GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, error)
		CheckEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error)
		Update(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error)
		Delete(ctx context.Context, tx *gorm.DB, userId string) error
		GetUserRoleById(ctx context.Context, tx *gorm.DB, userId string) (string, error)
		GetUserProjectsByUserId(ctx context.Context, tx *gorm.DB, userId string) ([]entity.Project, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Register(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetAllUserWithPagination(
	ctx context.Context,
	tx *gorm.DB,
	req dto.PaginationRequest,
) (dto.GetAllUserRepositoryResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var users []entity.User
	var err error
	var count int64

	req.Default()

	query := tx.WithContext(ctx).Model(&entity.User{})
	if req.Search != "" {
		query = query.Where("name LIKE ?", "%"+req.Search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return dto.GetAllUserRepositoryResponse{}, err
	}

	if err := query.Scopes(Paginate(req)).Find(&users).Error; err != nil {
		return dto.GetAllUserRepositoryResponse{}, err
	}

	totalPage := TotalPage(count, int64(req.PerPage))
	return dto.GetAllUserRepositoryResponse{
		Users: users,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}

func (r *userRepository) GetUserById(ctx context.Context, tx *gorm.DB, userId string) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	println("GetUserById in User Repo called with userId:", userId)

	var user entity.User
	if err := tx.WithContext(ctx).Where("user_id = ?", userId).Take(&user).Error; err != nil {
		log.Println("Error in GetUserById:", err)
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	println("GetUserByEmail called with email:", email)

	var user entity.User
	if err := tx.WithContext(ctx).Where("email = ?", email).Take(&user).Error; err != nil {
		print("Error in GetUserByEmail:", err)
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) CheckEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error) {
	if tx == nil {
		tx = r.db
	}

	var user entity.User
	if err := tx.WithContext(ctx).Where("email = ?", email).Take(&user).Error; err != nil {
		return entity.User{}, false, err
	}

	return user, true, nil
}

func (r *userRepository) Update(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) Delete(ctx context.Context, tx *gorm.DB, userId string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.User{}, "id = ?", userId).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUserRoleById(ctx context.Context, tx *gorm.DB, userId string) (string, error) {
	if tx == nil {
		tx = r.db
	}

	var user entity.User
	if err := tx.WithContext(ctx).Select("role").Where("id = ?", userId).Take(&user).Error; err != nil {
		return "", err
	}

	return user.Role, nil
}
func (r *userRepository) GetUserProjectsByUserId(ctx context.Context, tx *gorm.DB, userId string) ([]entity.Project, error) {
	if tx == nil {
		tx = r.db
	}

	var projects []entity.Project
	if err := tx.WithContext(ctx).Model(&entity.Project{}).
		Joins("JOIN project_members ON project_members.projects_project_id = projects.project_id").
		Where("project_members.users_user_id = ?", userId).
		Find(&projects).Error; err != nil {
		return nil, err
	}

	return projects, nil
}
