package provider

import (
	"fp_mbd/controller"
	"fp_mbd/repository"
	"fp_mbd/service"

	"github.com/minio/minio-go/v7"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func ProvideDocumentDependencies(injector *do.Injector, db *gorm.DB, minioClient *minio.Client) {
	// Repository
	documentRepository := repository.NewDocumentRepository(db)
	minioRepository := repository.NewMinioRepository(minioClient, "main")

	// Service
	documentService := service.NewDocumentService(documentRepository, minioRepository, db)

	// Controller
	do.Provide(
		injector, func(i *do.Injector) (controller.DocumentController, error) {
			return controller.NewDocumentController(documentService), nil
		},
	)
}
