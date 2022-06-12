package repository

import (
	"context"
	"database/sql"
	"myfin/entity/domain"
)

type GetSaldoRepository interface {
	GetSaldo(ctx context.Context, tx *sql.Tx, saldoId int) (domain.Saldo, error)
	MinSaldo(ctx context.Context, tx *sql.Tx, saldo domain.Saldo) domain.Saldo
}
