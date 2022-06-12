package service

import (
	"context"
	"myfin/entity/web"
)

type AvaUploadService interface {
	AvaUpload(ctx context.Context, request web.AvaUploadRequest) web.AvaUploadResponse
	FindById(ctx context.Context, nasabahId int) web.AvaUploadResponse
}
