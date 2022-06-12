package repository

import (
	"context"
	"database/sql"
	"errors"
	"myfin/entity/domain"
	"myfin/helper"
)

type GetNikRepositoryImpl struct {
}

func NewGetNikRepository() *GetNikRepositoryImpl {
	return &GetNikRepositoryImpl{}
}

func (repository *GetNikRepositoryImpl) GetNik(ctx context.Context, tx *sql.Tx, nik int) (domain.Nik, error) {
	SQL := "select nik, nama, alamat, tgl_lahir, kota_lahir from tb_nik where nik = ?"
	rows, err := tx.QueryContext(ctx, SQL, nik)
	helper.PanicIfError(err)
	defer rows.Close()

	nikNasabah := domain.Nik{}
	if rows.Next() {
		err := rows.Scan(&nikNasabah.Nik, &nikNasabah.Nama, &nikNasabah.Alamat, &nikNasabah.TanggalLahir, &nikNasabah.KotaLahir)
		helper.PanicIfError(err)
		return nikNasabah, nil
	} else {
		return nikNasabah, errors.New("Not Found")
	}
}
