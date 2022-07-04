package repository

import (
	"context"
	"database/sql"
	"errors"
	"myfin/entity/domain"
	"myfin/helper"
)

type TransaksiRepositoryImpl struct {
}

func NewTransaksiRepositoryImpl() *TransaksiRepositoryImpl {
	return &TransaksiRepositoryImpl{}
}

func (repository *TransaksiRepositoryImpl) FindTransaksiKredit(ctx context.Context, tx *sql.Tx, nasabah domain.TransaksiNasabah) (domain.TransaksiNasabah, error) {
	script := "select id_nasabah, id_pengajuan_kredit, tanggal_pengajuan, besar_pengajuan, bsr_pengajuan_diterima, status, ket_catatan from tb_pengajuan_kredit where id_nasabah = ? and status = ? limit 1"
	rows, err := tx.QueryContext(ctx, script, nasabah.IdUser, nasabah.Status)
	transaksinasabah := domain.TransaksiNasabah{}
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&transaksinasabah.IdTransaksi, &transaksinasabah.IdTransaksi, &transaksinasabah.Tanggal, &transaksinasabah.Nominal, &transaksinasabah.NominalDiterima, &transaksinasabah.Status, &transaksinasabah.Catatan)
		helper.PanicIfError(err)
		return transaksinasabah, nil
	} else {
		return transaksinasabah, errors.New("Transaksi Tidak Ditemukan")
	}
}

func (repository *TransaksiRepositoryImpl) FindTransaksiPenarikan(ctx context.Context, tx *sql.Tx, nasabah domain.TransaksiNasabah) (domain.TransaksiNasabah, error) {
	script := "select id_nasabah, id_pengajuan_kredit, tanngal_pengajuan, besar_pengajuan, status, ket_catatan from tb_pengajuan_kredit where id_nasabah = ? and status = ? limit 5"
	rows, err := tx.QueryContext(ctx, script, nasabah.IdUser, nasabah.Status)
	transaksinasabah := domain.TransaksiNasabah{}
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&transaksinasabah.IdUser, &transaksinasabah.IdTransaksi, &transaksinasabah.Tanggal, &transaksinasabah.Nominal, &transaksinasabah.Status, &transaksinasabah.Catatan)
		helper.PanicIfError(err)
		return transaksinasabah, nil
	} else {
		return transaksinasabah, errors.New("Transaksi Tidak Ditemukan")
	}
}
