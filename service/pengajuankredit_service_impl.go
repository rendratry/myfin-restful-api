package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"myfin/entity/domain"
	"myfin/entity/web"
	"myfin/helper"
	"myfin/repository"
	"time"
)

type PengajuanKreditServiceImpl struct {
	PengajuanKreditRepository repository.PengajuanKreditRepository
	DB                        *sql.DB
	Validate                  *validator.Validate
}

func NewPengajuanKreditService(pengajuankreditRepository repository.PengajuanKreditRepository, DB *sql.DB, validate *validator.Validate) *PengajuanKreditServiceImpl {
	return &PengajuanKreditServiceImpl{
		PengajuanKreditRepository: pengajuankreditRepository,
		DB:                        DB,
		Validate:                  validate}
}

func (service *PengajuanKreditServiceImpl) AjuanKredit(ctx context.Context, request web.PengajuanKreditRequest) web.PengajuanKreditResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dt := time.Now()
	layout := "2 Jan 2006 15:04"
	date := dt.Format(layout)

	pengajuankredit := domain.PengajuanKredit{
		IdNasabah:        request.IdNasabah,
		Penggunaan:       request.Penggunaan,
		BesarPengajuan:   request.BesarPengajuan,
		Tenor:            request.Tenor,
		TanggalPengajuan: date,
		Score:            request.Score,
		Status:           "pending",
	}
	pengajuankredit = service.PengajuanKreditRepository.AjuanKredit(ctx, tx, pengajuankredit)

	return helper.ToPengajuanKreditResponse(pengajuankredit)
}
