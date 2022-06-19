package repository

import (
	"context"
	"database/sql"
	"myfin/entity/domain"
	"myfin/helper"
)

type PengajuanKreditRepositoryImpl struct {
}

func NewPengajuanKreditRepository() *PengajuanKreditRepositoryImpl {
	return &PengajuanKreditRepositoryImpl{}
}

func (repository *PengajuanKreditRepositoryImpl) AjuanKredit(ctx context.Context, tx *sql.Tx, ajuankredit domain.PengajuanKredit) domain.PengajuanKredit {
	script := "insert into tb_pengajuan_kredit(id_nasabah, penggunaan, pekerjaan, gaji, gaji_tambahan, tanggal_pengajuan, besar_pengajuan, tenor, score, status) values (?,?,?,?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, script, ajuankredit.IdNasabah, ajuankredit.Penggunaan, ajuankredit.Pekerjaan, ajuankredit.Gaji, ajuankredit.GajiTambahan, ajuankredit.TanggalPengajuan, ajuankredit.BesarPengajuan, ajuankredit.Tenor, ajuankredit.Score, ajuankredit.Status)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	ajuankredit.IdNasabah = int(id)

	return ajuankredit
}
