package config

import (
	"fmt"
	"os"
	"strconv"

	"fp_mbd/constants"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func SetupMinioConnection() *minio.Client {
	if os.Getenv("APP_ENV") != constants.ENUM_RUN_PRODUCTION {
		err := godotenv.Load(".env")
		if err != nil {
			panic(err)
		}
	}

	host := os.Getenv("MINIO_HOST")
	port := os.Getenv("MINIO_PORT")
	access_key := os.Getenv("MINIO_ACCESS)KEY")
	secret_key := os.Getenv("MINIO_SECRET_KEY")
	useSSL, err := strconv.ParseBool(os.Getenv("MINIO_USE_SSL"))
	// bucket := os.Getenv("MINIO_BUCKET")

	if err != nil {
		panic(err)
	}

	endpoint := fmt.Sprintf("%s:%s", host, port)

	client, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(access_key, secret_key, ""),
		Secure: useSSL,
	})

	if err != nil {
		panic(err)
	}

	// set bucket policy ?

	return client
}