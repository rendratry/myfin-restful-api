package repository

import (
	"context"
	"database/sql"
	"errors"
	"myfin/entity/domain"
	"myfin/helper"
)

type DatanasabahRepositoryImpl struct {
}

func NewDatanasabahRepository() *DatanasabahRepositoryImpl {
	return &DatanasabahRepositoryImpl{}
}

func (repository *DatanasabahRepositoryImpl) UpdateDataNasabah(ctx context.Context, tx *sql.Tx, nasabah domain.DataNasabah) domain.DataNasabah {
	script := "update tb_data_nasabah set nik = ?, nama_lengkap = ?, no_hp = ?, alamat = ?, tgl_lahir = ?, kota_lahir = ?, ktp_file = ?, swa_file = ? where id_user = ?"
	_, err := tx.ExecContext(ctx, script, nasabah.Nik, nasabah.NamaLengkap, nasabah.NoHp, nasabah.Alamat, nasabah.TglLahir, nasabah.KotaLahir, nasabah.KtpFile, nasabah.SwaFile, nasabah.IdUser)
	helper.PanicIfError(err)
	return nasabah
}

func (repository *DatanasabahRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id_user int) (domain.DataNasabah, error) {
	script := "select id_user, nik, nama_lengkap, no_hp, alamat, tgl_lahir, kota_lahir,ktp_file, swa_file from tb_data_nasabah where id_user = ?"
	rows, err := tx.QueryContext(ctx, script, id_user)
	nasabah := domain.DataNasabah{}
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&nasabah.IdUser, &nasabah.Nik, &nasabah.NamaLengkap, &nasabah.NoHp, &nasabah.Alamat, &nasabah.TglLahir, &nasabah.KotaLahir, &nasabah.KtpFile, &nasabah.SwaFile)
		helper.PanicIfError(err)
		return nasabah, nil
	} else {
		return nasabah, errors.New("id Not Found")
	}
}
