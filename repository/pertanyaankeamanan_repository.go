package repository

import (
	"context"
	"database/sql"
	"myfin/entity/domain"
)

type KeamananNasabahRepository interface {
	UpdateKeamanan(ctx context.Context, tx *sql.Tx, nasabah domain.KeamananNasabah) domain.KeamananNasabah
	FindById(ctx context.Context, tx *sql.Tx, id_user int) (domain.KeamananNasabah, error)
}
