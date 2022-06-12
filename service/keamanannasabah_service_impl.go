package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
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

func (service *KeamananNasabahServiceImpl) FindById(ctx context.Context, datanasabahId int) web.KeamananNasabahResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	datanasabah, err := service.KeamananNasabahRepository.FindById(ctx, tx, datanasabahId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToKeamananNasabahResponse(datanasabah)
}
