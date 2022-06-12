package repository

import (
	"context"
	"database/sql"
	"myfin/entity/domain"
)

type GetNikRepository interface {
	GetNik(ctx context.Context, tx *sql.Tx, nik int) (domain.Nik, error)
}
