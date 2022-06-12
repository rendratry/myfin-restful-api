package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"myfin/entity/web"
	"myfin/helper"
	"myfin/repository"
)

type DatanasabahServiceImpl struct {
	DatanasabahRepository repository.DatanasabahRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewDatanasabahService(datanasabahRepository repository.DatanasabahRepository, DB *sql.DB, validate *validator.Validate) *DatanasabahServiceImpl {
	return &DatanasabahServiceImpl{
		DatanasabahRepository: datanasabahRepository,
		DB:                    DB,
		Validate:              validate}
}

func (service *DatanasabahServiceImpl) UpdateDataNasabah(ctx context.Context, request web.DatanasabahUpdateRequest) web.DataNasabahResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	datanasabah, err := service.DatanasabahRepository.FindById(ctx, tx, request.Id_user)
	helper.PanicIfError(err)

	datanasabah.Nik = request.Nik
	datanasabah.NamaLengkap = request.Nama_lengkap
	datanasabah.Alamat = request.Alamat
	datanasabah.NoHp = request.No_hp
	datanasabah.TglLahir = request.Tgl_lahir
	datanasabah.KotaLahir = request.Kota_lahir
	datanasabah.KtpFile = request.KtpFile
	datanasabah.SwaFile = request.SwaFile

	datanasabah = service.DatanasabahRepository.UpdateDataNasabah(ctx, tx, datanasabah)
	return helper.ToDataNasabahResponse(datanasabah)
}

func (service *DatanasabahServiceImpl) FindById(ctx context.Context, datanasabahId int) web.DataNasabahResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	helper.CommitOrRollback(tx)

	datanasabah, err := service.DatanasabahRepository.FindById(ctx, tx, datanasabahId)
	helper.PanicIfError(err)

	return helper.ToDataNasabahResponse(datanasabah)
}
