package service

import (
	"context"
	"myfin/entity/web"
)

type GetProfileService interface {
	GetProfile(ctx context.Context, id_user int) web.GetProfileResponse
}
