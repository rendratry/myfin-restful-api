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

type LoginNasabahServiceImpl struct {
	LoginNasabahRepository repository.LoginNasabahRepository
	DB                     *sql.DB
	Validate               *validator.Validate
}

func NewLoginNasabahService(loginnasabahRepository repository.LoginNasabahRepository, DB *sql.DB, validate *validator.Validate) *LoginNasabahServiceImpl {
	return &LoginNasabahServiceImpl{
		LoginNasabahRepository: loginnasabahRepository,
		DB:                     DB,
		Validate:               validate}
}

func (service *LoginNasabahServiceImpl) FindByEmail(ctx context.Context, request web.LoginNasabahRequest) web.LoginNasabahResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	email := domain.LoginNasabah{Email: request.Email}

	loginnasabah, err := service.LoginNasabahRepository.FindByEmail(ctx, tx, email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	_, err = helper.CheckPasswordHash(request.Pin, loginnasabah.Pin)
	helper.PanicIfError(err)

	loginnasabah.Email = request.Email
	loginnasabah.Status = "success"

	return helper.ToLoginNasabahResponse(loginnasabah)
}

//func (service *LoginNasabahServiceImpl) FindById(ctx context.Context, idUser int) web.LoginNasabahResponse {
//	tx, err := service.DB.Begin()
//	helper.PanicIfError(err)
//	defer helper.CommitOrRollback(tx)
//
//	loginnasabah, err := service.LoginNasabahRepository.FindById(ctx, tx, idUser)
//	if err != nil {
//		panic(exception.NewNotFoundError(err.Error()))
//	}
//
//	return helper.ToLoginNasabahResponse(loginnasabah)
//}
