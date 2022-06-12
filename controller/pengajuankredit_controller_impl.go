package controller

import (
	"github.com/julienschmidt/httprouter"
	"myfin/entity/web"
	"myfin/helper"
	"myfin/service"
	"net/http"
)

type PengajuanKreditControllerImpl struct {
	PengajuanKreditService service.PengajuanKreditService
}

func NewPengajuanKreditControllerImpl(pengajuankreditService service.PengajuanKreditService) *PengajuanKreditControllerImpl {
	return &PengajuanKreditControllerImpl{
		PengajuanKreditService: pengajuankreditService,
	}
}

func (controller *PengajuanKreditControllerImpl) AjuanKredit(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	pengajuankreditCreateRequest := web.PengajuanKreditRequest{}
	helper.ReadFromRequestBody(request, &pengajuankreditCreateRequest)

	pengajuankreditResponse := controller.PengajuanKreditService.AjuanKredit(request.Context(), pengajuankreditCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   pengajuankreditResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
