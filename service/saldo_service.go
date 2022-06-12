package service

import (
	"context"
	"myfin/entity/web"
)

type GetSaldoService interface {
	GetSaldo(ctx context.Context, saldoId int) web.GetSaldoResponse
	MinSaldo(ctx context.Context, request web.MinSaldoUpdateRequest) web.GetSaldoResponse
}
