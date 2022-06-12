package service

import (
	"context"
	"myfin/entity/web"
)

type LoginNasabahService interface {
	FindByEmail(ctx context.Context, request web.LoginNasabahRequest) web.LoginNasabahResponse
	//FindById(ctx context.Context, idUser int) web.LoginNasabahResponse
}
