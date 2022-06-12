package service

import (
	"context"
	"myfin/entity/web"
)

type PenarikanSaldoService interface {
	PenarikanSaldo(ctx context.Context, request web.PenarikanSaldoRequest) web.PenarikanSaldoResponse
}
