package repository

import (
	"context"
	"database/sql"
	"errors"
	"myfin/entity/domain"
	"myfin/helper"
)

type GetSaldoRepositoryImpl struct {
}

func NewGetSaldoRepository() *GetSaldoRepositoryImpl {
	return &GetSaldoRepositoryImpl{}
}

func (repository *GetSaldoRepositoryImpl) GetSaldo(ctx context.Context, tx *sql.Tx, saldoId int) (domain.Saldo, error) {
	SQL := "select id_user, saldo from tb_data_nasabah where id_user = ?"
	rows, err := tx.QueryContext(ctx, SQL, saldoId)
	helper.PanicIfError(err)
	defer rows.Close()

	saldo := domain.Saldo{}
	if rows.Next() {
		err := rows.Scan(&saldo.IdNasabah, &saldo.Saldo)
		helper.PanicIfError(err)
		return saldo, nil
	} else {
		return saldo, errors.New("Not Found")
	}
}

func (repository *GetSaldoRepositoryImpl) MinSaldo(ctx context.Context, tx *sql.Tx, saldo domain.Saldo) domain.Saldo {
	script := "update tb_data_nasabah set saldo = ? where id_user = ?"
	_, err := tx.ExecContext(ctx, script, saldo.Saldo, saldo.IdNasabah)
	helper.PanicIfError(err)
	return saldo
}
