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

type AvaUploadServiceImpl struct {
	AvaUploadRepository repository.AvaUploadRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewAvaUploadService(avauploadRepository repository.AvaUploadRepository, DB *sql.DB, validate *validator.Validate) *AvaUploadServiceImpl {
	return &AvaUploadServiceImpl{
		AvaUploadRepository: avauploadRepository,
		DB:                  DB,
		Validate:            validate}
}

func (service *AvaUploadServiceImpl) AvaUpload(ctx context.Context, request web.AvaUploadRequest) web.AvaUploadResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	avanasabah, err := service.AvaUploadRepository.FindById(ctx, tx, request.Id_user)
	helper.PanicIfError(err)

	avanasabah.Ava = request.Ava

	avanasabah = service.AvaUploadRepository.AvaUpload(ctx, tx, avanasabah)

	return helper.ToAvaUploadResponse(avanasabah)
}

func (service *AvaUploadServiceImpl) FindById(ctx context.Context, nasabahId int) web.AvaUploadResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	datanasabah, err := service.AvaUploadRepository.FindById(ctx, tx, nasabahId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToAvaUploadResponse(datanasabah)
}
