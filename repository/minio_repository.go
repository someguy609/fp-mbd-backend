package repository

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

type (
	MinioRepository interface {
		Upload(ctx context.Context, objectName string, header *multipart.FileHeader) (string, error)
		Get(ctx context.Context, objectName string) (string, error)
		Delete(ctx context.Context, objectName string) error
	}

	minioRepository struct {
		client *minio.Client
		bucketName string
	}
)

func NewMinioRepository(client *minio.Client, bucketName string) MinioRepository {
	return &minioRepository{
		client: client,
		bucketName: bucketName,
	}
}

func (r *minioRepository) Upload(ctx context.Context, objectName string, header *multipart.FileHeader) (string, error) {
	file, err := header.Open()

	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = r.client.PutObject(ctx, r.bucketName, objectName, file, header.Size, minio.PutObjectOptions{
		ContentType: header.Header.Get("Content-Type"),
	})

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("http://%s/%s/%s", r.client.EndpointURL().Host, r.bucketName, objectName), nil
}

func (r *minioRepository) Get(ctx context.Context, objectName string) (string, error) {
	_, err := r.client.StatObject(ctx, r.bucketName, objectName, minio.StatObjectOptions{})

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("http://%s/%s/%s", r.client.EndpointURL().Host, r.bucketName, objectName), nil
}

func (r *minioRepository) Delete(ctx context.Context, objectName string) error {
	err := r.client.RemoveObject(ctx, r.bucketName, objectName, minio.RemoveObjectOptions{})

	if err != nil {
		return err
	}
	
	return nil
}
