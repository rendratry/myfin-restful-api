package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"myfin/entity/domain"
	"myfin/entity/web"
	"myfin/helper"
	"myfin/repository"
)

type PenarikanSaldoServiceImpl struct {
	PenarikanSaldoRepository repository.PenarikanSaldoRepository
	DB                       *sql.DB
	Validate                 *validator.Validate
}

func NewPenarikanSaldoServiceImpl(penarikansaldoRepository repository.PenarikanSaldoRepository, DB *sql.DB, validate *validator.Validate) *PenarikanSaldoServiceImpl {
	return &PenarikanSaldoServiceImpl{
		PenarikanSaldoRepository: penarikansaldoRepository,
		DB:                       DB,
		Validate:                 validate,
	}
}

func (service *PenarikanSaldoServiceImpl) PenarikanSaldo(ctx context.Context, request web.PenarikanSaldoRequest) web.PenarikanSaldoResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	penarikansaldo := domain.PenarikanSaldo{
		JumlahPenarikan: request.JumlahPenarikan,
		Bank:            request.Bank,
		NoRek:           request.NoRek,
		NamaPemilik:     request.NamaPemilik,
		Status:          "verifikasi",
		IdNasabah:       request.IdNasabah,
	}
	penarikansaldo = service.PenarikanSaldoRepository.PenarikanSaldo(ctx, tx, penarikansaldo)

	return helper.ToPenarikanSaldoResponse(penarikansaldo)

}
