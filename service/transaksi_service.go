package service

import (
	"context"
	"myfin/entity/web"
)

type TransaksiService interface {
	FindTransaksiKredit(ctx context.Context, request web.TransaksiRequest) web.TransaksiResponse
}
