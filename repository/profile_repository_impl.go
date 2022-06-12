package repository

import (
	"context"
	"database/sql"
	"errors"
	"myfin/entity/domain"
	"myfin/helper"
)

type GetProfileRepositoryImpl struct {
}

func NewGetProfileRepository() *GetProfileRepositoryImpl {
	return &GetProfileRepositoryImpl{}
}

func (repository *GetProfileRepositoryImpl) GetProfile(ctx context.Context, tx *sql.Tx, id_user int) (domain.DataNasabah, error) {
	SQL := "select nama_lengkap, email, ava, no_hp, nik, alamat from tb_data_nasabah where id_user = ?"
	rows, err := tx.QueryContext(ctx, SQL, id_user)
	helper.PanicIfError(err)
	defer rows.Close()

	profileNasabah := domain.DataNasabah{}
	if rows.Next() {
		err := rows.Scan(&profileNasabah.NamaLengkap, &profileNasabah.Email, &profileNasabah.Ava, &profileNasabah.NoHp, &profileNasabah.Nik, &profileNasabah.Alamat)
		helper.PanicIfError(err)
		return profileNasabah, nil
	} else {
		return profileNasabah, errors.New("Not Found")
	}
}
