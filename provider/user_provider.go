package provider

import (
	"fp_mbd/controller"
	"fp_mbd/repository"
	"fp_mbd/service"

	"github.com/samber/do"
	"gorm.io/gorm"
)

func ProvideUserDependencies(injector *do.Injector, db *gorm.DB, jwtService service.JWTService) {
	// Repository
	userRepository := repository.NewUserRepository(db)
	refreshTokenRepository := repository.NewRefreshTokenRepository(db)

	// Service
	userService := service.NewUserService(userRepository, refreshTokenRepository, jwtService, db)

	// Controller
	do.Provide(
		injector, func(i *do.Injector) (controller.UserController, error) {
			return controller.NewUserController(userService), nil
		},
	)
}
