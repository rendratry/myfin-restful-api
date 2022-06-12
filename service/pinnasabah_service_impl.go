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

type PinNasabahServiceImpl struct {
	PinNasabahRepository repository.PinNasabahRepository
	DB                   *sql.DB
	Validate             *validator.Validate
}

func NewPinNasabahService(pinnasabahRepository repository.PinNasabahRepository, DB *sql.DB, validate *validator.Validate) *PinNasabahServiceImpl {
	return &PinNasabahServiceImpl{
		PinNasabahRepository: pinnasabahRepository,
		DB:                   DB,
		Validate:             validate}
}

func (service *PinNasabahServiceImpl) UpdatePin(ctx context.Context, request web.DatanasabahUpdatePinRequest) web.PinNasabahResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	pinnasabah, err := service.PinNasabahRepository.FindById(ctx, tx, request.Id_user)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	hashpassword, err := helper.HashPassword(request.Pin)
	helper.PanicIfError(err)

	pinnasabah.PinUser = hashpassword

	pinnasabah = service.PinNasabahRepository.UpdatePin(ctx, tx, pinnasabah)

	return helper.ToPinNasabahResponse(pinnasabah)
}

func (service *PinNasabahServiceImpl) FindById(ctx context.Context, datanasabahId int) web.PinNasabahResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	datanasabah, err := service.PinNasabahRepository.FindById(ctx, tx, datanasabahId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPinNasabahResponse(datanasabah)
}
