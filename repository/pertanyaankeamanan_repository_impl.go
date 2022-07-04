package repository

import (
	"context"
	"database/sql"
	"errors"
	"myfin/entity/domain"
	"myfin/helper"
)

type PertanyaanKeamananRepositoryImpl struct {
}

func NewPertanyaanKeamananRepository() *PertanyaanKeamananRepositoryImpl {
	return &PertanyaanKeamananRepositoryImpl{}
}

func (repository *PertanyaanKeamananRepositoryImpl) UpdateKeamanan(ctx context.Context, tx *sql.Tx, nasabah domain.KeamananNasabah) domain.KeamananNasabah {
	script := "update tb_data_nasabah set pertanyaan_keamanan = ?, jawaban_keamanan = ? where id_user = ?"
	_, err := tx.ExecContext(ctx, script, nasabah.PertanyaanKeamanan, nasabah.JawabanKeamanan, nasabah.Id)
	helper.PanicIfError(err)

	return nasabah
}

func (repository *PertanyaanKeamananRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id_user int) (domain.KeamananNasabah, error) {
	script := "select id_user, pertanyaan_keamanan, jawaban_keamanan from tb_data_nasabah where id_user = ?"
	rows, err := tx.QueryContext(ctx, script, id_user)
	nasabah := domain.KeamananNasabah{}
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&nasabah.Id, &nasabah.PertanyaanKeamanan, &nasabah.JawabanKeamanan)
		helper.PanicIfError(err)
		return nasabah, nil
	} else {
		return nasabah, errors.New("id Not Found")
	}
}

func (repository *PertanyaanKeamananRepositoryImpl) FindScurity(ctx context.Context, tx *sql.Tx, id_user domain.KeamananNasabah) (domain.KeamananNasabah, error) {
	script := "select id_user, email,pertanyaan_keamanan, jawaban_keamanan from tb_data_nasabah where id_user = ? and pertanyaan_keamanan = ? and jawaban_keamanan = ? limit 1"
	rows, err := tx.QueryContext(ctx, script, id_user.Id, id_user.PertanyaanKeamanan, id_user.JawabanKeamanan)
	nasabah := domain.KeamananNasabah{}
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&nasabah.Id, &nasabah.Email, &nasabah.PertanyaanKeamanan, &nasabah.JawabanKeamanan)
		helper.PanicIfError(err)
		return nasabah, nil
	} else {
		return nasabah, errors.New("id Not Found")
	}
}
