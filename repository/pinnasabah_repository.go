package repository

import (
	"context"
	"database/sql"
	"myfin/entity/domain"
)

type PinNasabahRepository interface {
	UpdatePin(ctx context.Context, tx *sql.Tx, nasabah domain.PinNasabah) domain.PinNasabah
	FindById(ctx context.Context, tx *sql.Tx, id_user int) (domain.PinNasabah, error)
}
