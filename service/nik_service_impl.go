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

type GetNikServiceImpl struct {
	GetNikRepository repository.GetNikRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewGetNikService(getnikRepository repository.GetNikRepository, DB *sql.DB, validate *validator.Validate) *GetNikServiceImpl {
	return &GetNikServiceImpl{
		GetNikRepository: getnikRepository,
		DB:               DB,
		Validate:         validate}
}

func (service *GetNikServiceImpl) GetNik(ctx context.Context, nik int) web.GetNikResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	nikNasabah, err := service.GetNikRepository.GetNik(ctx, tx, nik)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToGetNikResponse(nikNasabah)
}
