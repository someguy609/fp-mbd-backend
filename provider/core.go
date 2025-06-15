package provider

import (
	"fp_mbd/config"
	"fp_mbd/constants"
	"fp_mbd/service"

	"github.com/minio/minio-go/v7"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func InitDatabase(injector *do.Injector) {
	do.ProvideNamed(injector, constants.DB, func(i *do.Injector) (*gorm.DB, error) {
		return config.SetUpDatabaseConnection(), nil
	})
}

func InitMinioClient(injector *do.Injector) {
	do.ProvideNamed(injector, constants.Minio, func(i *do.Injector) (*minio.Client, error) {
		return config.SetupMinioConnection(), nil
	})
}

func RegisterDependencies(injector *do.Injector) {
	InitDatabase(injector)
	InitMinioClient(injector)

	do.ProvideNamed(injector, constants.JWTService, func(i *do.Injector) (service.JWTService, error) {
		return service.NewJWTService(), nil
	})

	// Initialize
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)
	minioClient := do.MustInvokeNamed[*minio.Client](injector, constants.Minio)

	// Provide Dependencies
	ProvideUserDependencies(injector, db, jwtService)
	ProvideDocumentDependencies(injector, db, minioClient)
	ProvideProjectDependencies(injector, db, minioClient)
}
