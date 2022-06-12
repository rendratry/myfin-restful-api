package repository

import (
	"context"
	"database/sql"
	"myfin/entity/domain"
)

type PenarikanSaldoRepository interface {
	PenarikanSaldo(ctx context.Context, tx *sql.Tx, saldo domain.PenarikanSaldo) domain.PenarikanSaldo
}
