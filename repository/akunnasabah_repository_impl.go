package repository

import (
	"context"
	"database/sql"
	"myfin/entity/domain"
	"myfin/helper"
)

type AkunNasabahRepositoryImpl struct {
}

func NewAkunNasabahRepositoryImpl() *AkunNasabahRepositoryImpl {
	return &AkunNasabahRepositoryImpl{}
}

func (repository *AkunNasabahRepositoryImpl) CreateAkun(ctx context.Context, tx *sql.Tx, akunnasabah domain.AkunNasabah) domain.AkunNasabah {
	script := "insert into tb_data_nasabah(email, kode_otp, pin_user, saldo, nik, nama_lengkap, no_hp, alamat, tgl_lahir, kota_lahir, ktp_file, swa_file, ava, pertanyaan_keamanan, jawaban_keamanan) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, script, akunnasabah.Email, akunnasabah.KodeOtp, akunnasabah.PinUser, akunnasabah.Saldo, akunnasabah.Nik, akunnasabah.NamaLengkap, akunnasabah.NoHp, akunnasabah.Alamat, akunnasabah.TglLahir, akunnasabah.KotaLahir, akunnasabah.KtpFile, akunnasabah.SwaFile, akunnasabah.Ava, akunnasabah.PertanyaanKeamanan, akunnasabah.JawabanKeamanan)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	akunnasabah.IdUser = int(id)

	return akunnasabah
}

//func (repository *AkunNasabahRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, emailUser string) (domain.AkunNasabah, error) {
//	script := "select email from tb_data_nasabah where email = ?"
//	rows, err := tx.QueryContext(ctx, script, emailUser)
//	emailcheck := domain.AkunNasabah{}
//	helper.PanicIfError(err)
//	defer rows.Close()
//
//	if rows.Next() {
//		err := rows.Scan(&emailcheck.Email)
//		helper.PanicIfError(err)
//		return emailcheck, nil
//	} else {
//		return emailcheck, errors.New("Email Belum Terdaftar")
//	}
//}
