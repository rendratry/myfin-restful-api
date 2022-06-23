package app

import (
	"github.com/julienschmidt/httprouter"
	"myfin/controller"
	"myfin/exception"
)

func NewRouter(
	datanasabahController controller.DatanasabahController,
	keamanannasabahController controller.KeamananNasabahController,
	akunnasabahController controller.AkunNasabahController,
	pengajuankreditController controller.PengajuanKreditController,
	emailvalidationController controller.EmailValidationController,
	penarikansaldoController controller.PenarikanSaldoController,
	getsaldoController controller.GetSaldoController,
	getnikController controller.GetNikController,
	loginController controller.LoginNasabahController,
	profileController controller.GetProfileController,
	avaController controller.AvaUploadController,
	transaksiController controller.TransaksiController,
	pinnasabahController controller.PinNasabahController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/transaksi/:nasabahId/:status", transaksiController.FindTransaksiKredit)
	router.POST("/api/keamanancheck/:datanasabahId", keamanannasabahController.FindById)
	router.POST("/api/nikcheck", akunnasabahController.FindByNik)
	router.POST("/api/emailcheck", akunnasabahController.FindByEmail)
	router.PUT("/api/ava/:nasabahId", avaController.AvaUpload)
	router.GET("/api/profile/:datanasabahId", profileController.GetProfile)
	router.POST("/api/getnik/", getnikController.GetNik)
	router.POST("/api/login", loginController.FindByEmail)
	router.PUT("/api/minsaldo/:saldoId", getsaldoController.MinSaldo)
	router.GET("/api/getsaldo/:saldoId", getsaldoController.GetSaldo)
	router.POST("/api/penarikansaldo", penarikansaldoController.PenarikanSaldo)
	router.GET("/api/datanasabah/:id_user", datanasabahController.FindById)
	router.POST("/api/pengajuankredit", pengajuankreditController.AjuanKredit)
	router.PUT("/api/datanasabah/:datanasabahId", datanasabahController.UpdateDataNasabah)
	router.POST("/api/datanasabah", akunnasabahController.CreateAkun)
	router.PUT("/api/pinnasabah/:datanasabahId", pinnasabahController.UpdatePin)
	router.PUT("/api/keamanannasabah/:datanasabahId", keamanannasabahController.UpdateKeamanan)
	router.POST("/api/emailvalidation", emailvalidationController.EmailValidation)

	router.PanicHandler = exception.ErrorHandler

	return router
}
