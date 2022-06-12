package repository

import (
	"context"
	"database/sql"
	"errors"
	"myfin/entity/domain"
	"myfin/helper"
)

type EmailValidationRepositoryImpl struct {
}

func NewEmailValidationRepository() *EmailValidationRepositoryImpl {
	return &EmailValidationRepositoryImpl{}
}

func (repository *EmailValidationRepositoryImpl) EmailValidation(ctx context.Context, tx *sql.Tx, validation domain.EmailValidation) (domain.EmailValidation, error) {
	script := "select email from tb_data_nasabah where email = ? and kode_otp = ? limit 1"
	rows, err := tx.QueryContext(ctx, script, validation.Email, validation.KodeOtp)
	emailvalidation := domain.EmailValidation{}
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&emailvalidation.Email)
		helper.PanicIfError(err)
		return emailvalidation, nil
	} else {
		return emailvalidation, errors.New("Kode OTP salah")
	}
}

func (repository *EmailValidationRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, idUser int) (domain.EmailValidation, error) {
	script := "select email, kode_otp from tb_data_nasabah where email = ?"
	rows, err := tx.QueryContext(ctx, script, idUser)
	emailvalidation := domain.EmailValidation{}
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&emailvalidation.Email)
		helper.PanicIfError(err)
		return emailvalidation, nil
	} else {
		return emailvalidation, errors.New("Email Belum Terdaftar")
	}
}
