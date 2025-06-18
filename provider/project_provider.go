package provider

import (
	"fp_mbd/controller"
	"fp_mbd/repository"
	"fp_mbd/service"

	"github.com/minio/minio-go/v7"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func ProvideProjectDependencies(injector *do.Injector, db *gorm.DB, minioClient *minio.Client) {
	// Repository
	userRepository := repository.NewUserRepository(db)
	projectRepository := repository.NewProjectRepository(db)
	projectMemberRepository := repository.NewProjectMemberRepository(db)
	documentRepository := repository.NewDocumentRepository(db)
	minioRepository := repository.NewMinioRepository(minioClient, "main")

	// Service
	projectService := service.NewProjectService(userRepository, projectRepository, documentRepository, projectMemberRepository, db)
	documentService := service.NewDocumentService(documentRepository, minioRepository, db)

	// Controller
	do.Provide(
		injector, func(i *do.Injector) (controller.ProjectController, error) {
			return controller.NewProjectController(projectService, documentService), nil
		},
	)
}
