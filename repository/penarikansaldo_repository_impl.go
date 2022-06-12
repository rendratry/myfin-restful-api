package repository

import (
	"context"
	"database/sql"
	"myfin/entity/domain"
	"myfin/helper"
)

type PenarikanSaldoRepositoryImpl struct {
}

func NewPenarikanSaldoRepositoryImpl() *PenarikanSaldoRepositoryImpl {
	return &PenarikanSaldoRepositoryImpl{}
}

func (repository *PenarikanSaldoRepositoryImpl) PenarikanSaldo(ctx context.Context, tx *sql.Tx, saldo domain.PenarikanSaldo) domain.PenarikanSaldo {
	script := "insert into tb_penarikan_saldo(jml_penarikan, bank, no_rek, nama_pemilik, status, id_nasabah) values (?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, script, saldo.JumlahPenarikan, saldo.Bank, saldo.NoRek, saldo.NamaPemilik, saldo.Status, saldo.IdNasabah)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	saldo.IdNasabah = int(id)

	return saldo
}
