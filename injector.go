//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"myfin/app"
	"myfin/controller"
	"myfin/middleware"
	"myfin/repository"
	"myfin/service"
	"net/http"
)

var datanasabahSet = wire.NewSet(
	repository.NewDatanasabahRepository,
	wire.Bind(new(repository.DatanasabahRepository), new(*repository.DatanasabahRepositoryImpl)),
	service.NewDatanasabahService,
	wire.Bind(new(service.DataNasabahService), new(*service.DatanasabahServiceImpl)),
	controller.NewDatanasabahController,
	wire.Bind(new(controller.DatanasabahController), new(*controller.DatanasabahControllerImpl)),
)
var keamanannasabahSet = wire.NewSet(
	repository.NewPertanyaanKeamananRepository,
	wire.Bind(new(repository.KeamananNasabahRepository), new(*repository.PertanyaanKeamananRepositoryImpl)),
	service.NewKeamananNasabahService,
	wire.Bind(new(service.KeamananNasabahService), new(*service.KeamananNasabahServiceImpl)),
	controller.NewKeamananNasabahController,
	wire.Bind(new(controller.KeamananNasabahController), new(*controller.KeamananNasabahControllerImpl)),
)

var emailvalidationSet = wire.NewSet(
	repository.NewEmailValidationRepository,
	wire.Bind(new(repository.EmailValidationRepository), new(*repository.EmailValidationRepositoryImpl)),
	service.NewEmailValidationService,
	wire.Bind(new(service.EmailValidationService), new(*service.EmailValidationServiceImpl)),
	controller.NewEmailValidationController,
	wire.Bind(new(controller.EmailValidationController), new(*controller.EmailValidationControllerImpl)),
)

var pinnasabahSet = wire.NewSet(
	repository.NewPinNasabahRepositoryImpl,
	wire.Bind(new(repository.PinNasabahRepository), new(*repository.PinNasabahRepositoryImpl)),
	service.NewPinNasabahService,
	wire.Bind(new(service.PinNasabahService), new(*service.PinNasabahServiceImpl)),
	controller.NewPinNasabahController,
	wire.Bind(new(controller.PinNasabahController), new(*controller.PinNasabahControllerImpl)),
)

var pengajuankreditSet = wire.NewSet(
	repository.NewPengajuanKreditRepository,
	wire.Bind(new(repository.PengajuanKreditRepository), new(*repository.PengajuanKreditRepositoryImpl)),
	service.NewPengajuanKreditService,
	wire.Bind(new(service.PengajuanKreditService), new(*service.PengajuanKreditServiceImpl)),
	controller.NewPengajuanKreditControllerImpl,
	wire.Bind(new(controller.PengajuanKreditController), new(*controller.PengajuanKreditControllerImpl)),
)

var penarikansaldoSet = wire.NewSet(
	repository.NewPenarikanSaldoRepositoryImpl,
	wire.Bind(new(repository.PenarikanSaldoRepository), new(*repository.PenarikanSaldoRepositoryImpl)),
	service.NewPenarikanSaldoServiceImpl,
	wire.Bind(new(service.PenarikanSaldoService), new(*service.PenarikanSaldoServiceImpl)),
	controller.NewPenarikanSaldoControllerImpl,
	wire.Bind(new(controller.PenarikanSaldoController), new(*controller.PenarikanSaldoControllerImpl)),
)

var akunnasabahSet = wire.NewSet(
	repository.NewAkunNasabahRepositoryImpl,
	wire.Bind(new(repository.AkunNasabahRepository), new(*repository.AkunNasabahRepositoryImpl)),
	service.NewAkunNasabahService,
	wire.Bind(new(service.AkunNasabahService), new(*service.AkunNasabahServiceImpl)),
	controller.NewAkunNasabahController,
	wire.Bind(new(controller.AkunNasabahController), new(*controller.AkunNasabahControllerImpl)),
)

var getsaldoSet = wire.NewSet(
	repository.NewGetSaldoRepository,
	wire.Bind(new(repository.GetSaldoRepository), new(*repository.GetSaldoRepositoryImpl)),
	service.NewGetSaldoService,
	wire.Bind(new(service.GetSaldoService), new(*service.GetSaldoServiceImpl)),
	controller.NewGetSaldoController,
	wire.Bind(new(controller.GetSaldoController), new(*controller.GetSaldoControllerImpl)),
)

var getnikSet = wire.NewSet(
	repository.NewGetNikRepository,
	wire.Bind(new(repository.GetNikRepository), new(*repository.GetNikRepositoryImpl)),
	service.NewGetNikService,
	wire.Bind(new(service.GetNikService), new(*service.GetNikServiceImpl)),
	controller.NewGetNikController,
	wire.Bind(new(controller.GetNikController), new(*controller.GetNikControllerImpl)),
)

var loginSet = wire.NewSet(
	repository.NewLoginNasabahRepository,
	wire.Bind(new(repository.LoginNasabahRepository), new(*repository.LoginNasabahRepositoryImpl)),
	service.NewLoginNasabahService,
	wire.Bind(new(service.LoginNasabahService), new(*service.LoginNasabahServiceImpl)),
	controller.NewLoginNasabahController,
	wire.Bind(new(controller.LoginNasabahController), new(*controller.LoginNasabahControllerImpl)),
)

var profileSet = wire.NewSet(
	repository.NewGetProfileRepository,
	wire.Bind(new(repository.GetProfileRepository), new(*repository.GetProfileRepositoryImpl)),
	service.NewGetProfileService,
	wire.Bind(new(service.GetProfileService), new(*service.GetProfileServiceImpl)),
	controller.NewGetProfileController,
	wire.Bind(new(controller.GetProfileController), new(*controller.GetProfileControllerImpl)),
)

var avaSet = wire.NewSet(
	repository.NewAvaUploadRepository,
	wire.Bind(new(repository.AvaUploadRepository), new(*repository.AvaUploadRepositoryImpl)),
	service.NewAvaUploadService,
	wire.Bind(new(service.AvaUploadService), new(*service.AvaUploadServiceImpl)),
	controller.NewAvaUploadController,
	wire.Bind(new(controller.AvaUploadController), new(*controller.AvaUploadControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.GetConnection,
		validator.New,
		profileSet,
		datanasabahSet,
		avaSet,
		loginSet,
		keamanannasabahSet,
		pinnasabahSet,
		akunnasabahSet,
		emailvalidationSet,
		pengajuankreditSet,
		penarikansaldoSet,
		getsaldoSet,
		getnikSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
