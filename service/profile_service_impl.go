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

type GetProfileServiceImpl struct {
	GetProfileRepository repository.GetProfileRepository
	DB                   *sql.DB
	Validate             *validator.Validate
}

func NewGetProfileService(getprofileRepository repository.GetProfileRepository, DB *sql.DB, validate *validator.Validate) *GetProfileServiceImpl {
	return &GetProfileServiceImpl{
		GetProfileRepository: getprofileRepository,
		DB:                   DB,
		Validate:             validate}
}

func (service *GetProfileServiceImpl) GetProfile(ctx context.Context, id_user int) web.GetProfileResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	profileNasabah, err := service.GetProfileRepository.GetProfile(ctx, tx, id_user)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToGetProfileResponse(profileNasabah)
}
