package service

import (
	"context"
	"myfin/entity/web"
)

type EmailValidationService interface {
	EmailValidation(ctx context.Context, request web.EmailValidationRequest) web.EmailValidationResponse
	FindById(ctx context.Context, idUser int) web.EmailValidationResponse
}
