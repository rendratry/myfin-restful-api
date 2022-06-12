package repository

import (
	"context"
	"database/sql"
	"errors"
	"myfin/entity/domain"
	"myfin/helper"
)

type AvaUploadRepositoryImpl struct {
}

func NewAvaUploadRepository() *AvaUploadRepositoryImpl {
	return &AvaUploadRepositoryImpl{}
}

func (repository *AvaUploadRepositoryImpl) AvaUpload(ctx context.Context, tx *sql.Tx, ava domain.Ava) domain.Ava {
	script := "update tb_data_nasabah set ava = ? where id_user = ?"
	_, err := tx.ExecContext(ctx, script, ava.Ava, ava.Id)
	helper.PanicIfError(err)
	return ava
}

func (repository *AvaUploadRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id_user int) (domain.Ava, error) {

	script := "select id_user, ava from tb_data_nasabah where id_user = ? limit 1"
	rows, err := tx.QueryContext(ctx, script, id_user)
	nasabah := domain.Ava{}
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&nasabah.Id, &nasabah.Ava)
		helper.PanicIfError(err)
		return nasabah, nil
	} else {
		return nasabah, errors.New("id Not Found")
	}
}
