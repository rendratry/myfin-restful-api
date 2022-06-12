package service

import (
	"context"
	"myfin/entity/web"
)

type GetNikService interface {
	GetNik(ctx context.Context, nik int) web.GetNikResponse
}
