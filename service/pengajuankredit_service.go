package service

import (
	"context"
	"myfin/entity/web"
)

type PengajuanKreditService interface {
	AjuanKredit(ctx context.Context, request web.PengajuanKreditRequest) web.PengajuanKreditResponse
}
