package repository

import (
	"context"
	"database/sql"
	"myfin/entity/domain"
)

type GetProfileRepository interface {
	GetProfile(ctx context.Context, tx *sql.Tx, id_user int) (domain.DataNasabah, error)
}
