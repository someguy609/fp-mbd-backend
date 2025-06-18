package provider

import (
	"fp_mbd/controller"
	"fp_mbd/repository"
	"fp_mbd/service"

	"github.com/samber/do"
	"gorm.io/gorm"
)

func ProvideProjectMemberDependencies(injector *do.Injector, db *gorm.DB) {
	// Repository
	userRepository := repository.NewUserRepository(db)
	projectMemberRepository := repository.NewProjectMemberRepository(db)
	// refreshTokenRepository := repository.NewRefreshTokenRepository(db)

	// Service
	projectMemberService := service.NewProjectMemberService(userRepository, projectMemberRepository, db)

	// Controller
	do.Provide(
		injector, func(i *do.Injector) (controller.ProjectMemberController, error) {
			return controller.NewProjectMemberController(projectMemberService), nil
		},
	)
}
