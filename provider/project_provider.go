package provider

import (
	"fp_mbd/controller"
	"fp_mbd/repository"
	"fp_mbd/service"

	"github.com/samber/do"
	"gorm.io/gorm"
)

func ProvideProjectDependencies(injector *do.Injector, db *gorm.DB) {
	// Repository
	userRepository := repository.NewUserRepository(db)
	projectRepository := repository.NewProjectRepository(db)

	// Service
	projectService := service.NewProjectService(userRepository, projectRepository, db)

	// Controller
	do.Provide(
		injector, func(i *do.Injector) (controller.ProjectController, error) {
			return controller.NewProjectController(projectService), nil
		},
	)
}
