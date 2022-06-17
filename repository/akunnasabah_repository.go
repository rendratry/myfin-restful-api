package repository

import (
	"context"
	"database/sql"
	"myfin/entity/domain"
)

type AkunNasabahRepository interface {
	CreateAkun(ctx context.Context, tx *sql.Tx, akunnasabah domain.AkunNasabah) domain.AkunNasabah
	FindByEmail(ctx context.Context, tx *sql.Tx, nasabah string) (domain.LoginNasabah, error)
	FindByNik(ctx context.Context, tx *sql.Tx, nik string) (domain.DataNasabah, error)
}
