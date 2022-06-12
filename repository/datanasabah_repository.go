package repository

import (
	"context"
	"database/sql"
	"myfin/entity/domain"
)

type DatanasabahRepository interface {
	UpdateDataNasabah(ctx context.Context, tx *sql.Tx, nasabah domain.DataNasabah) domain.DataNasabah
	FindById(ctx context.Context, tx *sql.Tx, id_user int) (domain.DataNasabah, error)
}
