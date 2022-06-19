package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"myfin/entity/domain"
	"myfin/entity/web"
	"myfin/exception"
	"myfin/helper"
	"myfin/repository"
)

type KeamananNasabahServiceImpl struct {
	KeamananNasabahRepository repository.KeamananNasabahRepository
	DB                        *sql.DB
	Validate                  *validator.Validate
}

func NewKeamananNasabahService(datanasabahRepository repository.KeamananNasabahRepository, DB *sql.DB, validate *validator.Validate) *KeamananNasabahServiceImpl {
	return &KeamananNasabahServiceImpl{
		KeamananNasabahRepository: datanasabahRepository,
		DB:                        DB,
		Validate:                  validate}
}

func (service *KeamananNasabahServiceImpl) UpdateKeamanan(ctx context.Context, request web.DatanasabahUpdateKeamananRequest) web.KeamananNasabahResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	datanasabah, err := service.KeamananNasabahRepository.FindById(ctx, tx, request.Id_user)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	datanasabah.PertanyaanKeamanan = request.Pertanyaan_Keamanan
	datanasabah.JawabanKeamanan = request.Jawaban_Keamanan

	datanasabah = service.KeamananNasabahRepository.UpdateKeamanan(ctx, tx, datanasabah)

	return helper.ToKeamananNasabahResponse(datanasabah)
}

func (service *KeamananNasabahServiceImpl) FindById(ctx context.Context, datanasabahId web.DatanasabahUpdateKeamananRequest) web.KeamananNasabahResponse {
	err := service.Validate.Struct(datanasabahId)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	datanasabah, err := service.KeamananNasabahRepository.FindById(ctx, tx, datanasabahId.Id_user)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	datanasabah.PertanyaanKeamanan = datanasabahId.Pertanyaan_Keamanan
	datanasabah.JawabanKeamanan = datanasabahId.Jawaban_Keamanan

	return helper.ToKeamananNasabahResponse(datanasabah)
}

func (service *KeamananNasabahServiceImpl) FindScurity(ctx context.Context, datanasabahId web.DatanasabahUpdateKeamananRequest) web.KeamananNasabahResponse {
	err := service.Validate.Struct(datanasabahId)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	keamanan := domain.KeamananNasabah{Id: datanasabahId.Id_user, PertanyaanKeamanan: datanasabahId.Pertanyaan_Keamanan, JawabanKeamanan: datanasabahId.Jawaban_Keamanan}

	datanasabah, err := service.KeamananNasabahRepository.FindScurity(ctx, tx, keamanan)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	datanasabah.PertanyaanKeamanan = datanasabahId.Pertanyaan_Keamanan
	datanasabah.JawabanKeamanan = datanasabahId.Jawaban_Keamanan

	return helper.ToKeamananNasabahResponse(datanasabah)
}