package repository

import (
	"context"
	"database/sql"
	"errors"
	"myfin/entity/domain"
	"myfin/helper"
)

type PinNasabahRepositoryImpl struct {
}

func NewPinNasabahRepositoryImpl() *PinNasabahRepositoryImpl {
	return &PinNasabahRepositoryImpl{}
}

func (repository *PinNasabahRepositoryImpl) UpdatePin(ctx context.Context, tx *sql.Tx, nasabah domain.PinNasabah) domain.PinNasabah {
	script := "update tb_data_nasabah set pin_user = ? where id_user = ?"
	_, err := tx.ExecContext(ctx, script, nasabah.PinUser, nasabah.IdUser)
	helper.PanicIfError(err)

	return nasabah
}

func (repository *PinNasabahRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id_user int) (domain.PinNasabah, error) {
	script := "select id_user, pin_user from tb_data_nasabah where id_user = ? "
	rows, err := tx.QueryContext(ctx, script, id_user)
	helper.PanicIfError(err)
	defer rows.Close()

	nasabah := domain.PinNasabah{}
	if rows.Next() {
		err := rows.Scan(&nasabah.IdUser, &nasabah.PinUser)
		helper.PanicIfError(err)
		return nasabah, nil
	} else {
		return nasabah, errors.New("id Not Found")
	}
}
