package service

import (
	"context"
	"myfin/entity/web"
)

type AkunNasabahService interface {
	CreateAkun(ctx context.Context, request web.AkunNasabahCreateRequest) web.AkunNasabahResponse
	//FindByEmail(ctx context.Context, emailUser string) web.EmailValidationResponse
}
