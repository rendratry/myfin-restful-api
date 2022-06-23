package repository

import (
	"context"
	"database/sql"
	"myfin/entity/domain"
)

type TransaksiRepository interface {
	FindTransaksiKredit(ctx context.Context, tx *sql.Tx, nasabah domain.TransaksiNasabah) (domain.TransaksiNasabah, error)
	FindTransaksiPenarikan(ctx context.Context, tx *sql.Tx, nasabah domain.TransaksiNasabah) (domain.TransaksiNasabah, error)
}
