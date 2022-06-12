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

type GetSaldoServiceImpl struct {
	GetSaldoRepository repository.GetSaldoRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewGetSaldoService(getsaldoRepository repository.GetSaldoRepository, DB *sql.DB, validate *validator.Validate) *GetSaldoServiceImpl {
	return &GetSaldoServiceImpl{
		GetSaldoRepository: getsaldoRepository,
		DB:                 DB,
		Validate:           validate}
}

func (service *GetSaldoServiceImpl) GetSaldo(ctx context.Context, saldoId int) web.GetSaldoResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	saldo, err := service.GetSaldoRepository.GetSaldo(ctx, tx, saldoId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToGetSaldoResponse(saldo)
}

func (service *GetSaldoServiceImpl) MinSaldo(ctx context.Context, request web.MinSaldoUpdateRequest) web.GetSaldoResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	minsaldo, err := service.GetSaldoRepository.GetSaldo(ctx, tx, request.IdNasabah)
	helper.PanicIfError(err)

	saldoint := request.Saldo
	minsaldoint := minsaldo.Saldo

	minsaldo.Saldo = minsaldoint - saldoint

	minsaldo = service.GetSaldoRepository.MinSaldo(ctx, tx, minsaldo)
	return helper.ToGetSaldoResponse(minsaldo)
}
