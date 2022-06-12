package service

import (
	"context"
	"myfin/entity/web"
)

type PinNasabahService interface {
	UpdatePin(ctx context.Context, request web.DatanasabahUpdatePinRequest) web.PinNasabahResponse
	FindById(ctx context.Context, datanasabahId int) web.PinNasabahResponse
}
