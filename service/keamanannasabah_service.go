package service

import (
	"context"
	"myfin/entity/web"
)

type KeamananNasabahService interface {
	UpdateKeamanan(ctx context.Context, request web.DatanasabahUpdateKeamananRequest) web.KeamananNasabahResponse
	FindById(ctx context.Context, datanasabahId web.DatanasabahUpdateKeamananRequest) web.KeamananNasabahResponse
	FindScurity(ctx context.Context, datanasabahId web.DatanasabahUpdateKeamananRequest) web.KeamananNasabahResponse
}
