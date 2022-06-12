package repository

import (
	"context"
	"database/sql"
	"myfin/entity/domain"
)

type AkunNasabahRepository interface {
	CreateAkun(ctx context.Context, tx *sql.Tx, akunnasabah domain.AkunNasabah) domain.AkunNasabah
	//FindByEmail(ctx context.Context, tx *sql.Tx, emailUser string) (domain.AkunNasabah, error)
}
