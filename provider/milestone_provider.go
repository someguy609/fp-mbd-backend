package provider

import (
	"fp_mbd/controller"
	"fp_mbd/repository"
	"fp_mbd/service"

	"github.com/samber/do"
	"gorm.io/gorm"
)

func ProvideMilestoneDependencies(injector *do.Injector, db *gorm.DB, jwtService service.JWTService) {
	// Repository
	userRepository := repository.NewUserRepository(db)
	milestoneRepository := repository.NewMilestoneRepository(db)
	projectMemberRepository := repository.NewProjectMemberRepository(db)
	// refreshTokenRepository := repository.NewRefreshTokenRepository(db)

	// Service
	milestoneService := service.NewMilestoneService(milestoneRepository, userRepository, projectMemberRepository, jwtService, db)

	// Controller
	do.Provide(
		injector, func(i *do.Injector) (controller.MilestoneController, error) {
			return controller.NewMilestoneController(milestoneService), nil
		},
	)
}
