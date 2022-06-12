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

type EmailValidationServiceImpl struct {
	EmailValidationRepository repository.EmailValidationRepository
	DB                        *sql.DB
	Validate                  *validator.Validate
}

func NewEmailValidationService(emailvalidationRepository repository.EmailValidationRepository, DB *sql.DB, validate *validator.Validate) *EmailValidationServiceImpl {
	return &EmailValidationServiceImpl{
		EmailValidationRepository: emailvalidationRepository,
		DB:                        DB,
		Validate:                  validate}
}

func (service *EmailValidationServiceImpl) EmailValidation(ctx context.Context, request web.EmailValidationRequest) web.EmailValidationResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	email := domain.EmailValidation{Email: request.Email, KodeOtp: request.KodeOtp}

	emailvalidation, err := service.EmailValidationRepository.EmailValidation(ctx, tx, email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	emailvalidation.Email = request.Email
	emailvalidation.KodeOtp = request.KodeOtp
	emailvalidation.Status = "Validated"

	return helper.ToEmailValidationResponse(emailvalidation)
}

func (service *EmailValidationServiceImpl) FindById(ctx context.Context, idUser int) web.EmailValidationResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	emailvalidation, err := service.EmailValidationRepository.FindById(ctx, tx, idUser)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToEmailValidationResponse(emailvalidation)
}
