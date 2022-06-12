package service

import (
	"context"
	"myfin/entity/web"
)

type DataNasabahService interface {
	UpdateDataNasabah(ctx context.Context, request web.DatanasabahUpdateRequest) web.DataNasabahResponse
	FindById(ctx context.Context, datanasabahId int) web.DataNasabahResponse
}
