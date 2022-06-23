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

type TransaksiServiceImpl struct {
	TransaksiRepository repository.TransaksiRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewTransaksiService(transaksiRepository repository.TransaksiRepository, DB *sql.DB, validate *validator.Validate) *TransaksiServiceImpl {
	return &TransaksiServiceImpl{
		TransaksiRepository: transaksiRepository,
		DB:                  DB,
		Validate:            validate}
}

func (service *TransaksiServiceImpl) FindTransaksiKredit(ctx context.Context, request web.TransaksiRequest) web.TransaksiResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	reqTransaksi := domain.TransaksiNasabah{IdUser: request.IdNasabah, Status: request.Status}

	transaksinasabah, err := service.TransaksiRepository.FindTransaksiKredit(ctx, tx, reqTransaksi)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToTransaksiResponse(transaksinasabah)
}
