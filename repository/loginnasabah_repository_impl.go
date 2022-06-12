package repository

import (
	"context"
	"database/sql"
	"errors"
	"myfin/entity/domain"
	"myfin/helper"
)

type LoginNasabahRepositoryImpl struct {
}

func NewLoginNasabahRepository() *LoginNasabahRepositoryImpl {
	return &LoginNasabahRepositoryImpl{}
}

func (repository *LoginNasabahRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, nasabah domain.LoginNasabah) (domain.LoginNasabah, error) {
	script := "select id_user, email, pin_user from tb_data_nasabah where email = ? limit 1"
	rows, err := tx.QueryContext(ctx, script, nasabah.Email)
	helper.PanicIfError(err)
	nasabah = domain.LoginNasabah{}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&nasabah.IdNasabah, &nasabah.Email, &nasabah.Pin)
		helper.PanicIfError(err)
		return nasabah, nil
	} else {
		return nasabah, errors.New("Email atau Pin Salah")
	}

}

//func (repository *LoginNasabahRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, idUser int) (domain.LoginNasabah, error) {
//	script := "select id_user, email, pin_user from tb_data_nasabah where email = ?"
//	rows, err := tx.QueryContext(ctx, script, idUser)
//	loginnasabah := domain.LoginNasabah{}
//	helper.PanicIfError(err)
//	defer rows.Close()
//
//	if rows.Next() {
//		err := rows.Scan(&loginnasabah.Email)
//		helper.PanicIfError(err)
//		return loginnasabah, nil
//	} else {
//		return loginnasabah, errors.New("Email atau Pin Salah")
//	}
//}
