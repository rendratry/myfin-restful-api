package service

import (
	"context"
	"myfin/entity/web"
)

type AkunNasabahService interface {
	CreateAkun(ctx context.Context, request web.AkunNasabahCreateRequest) web.AkunNasabahResponse
	FindByEmail(ctx context.Context, request web.LoginNasabahRequest) web.LoginNasabahResponse
	FindByNik(ctx context.Context, request web.CheckNikRequest) web.DataNasabahResponse
}
