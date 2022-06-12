package controller

import (
	"github.com/julienschmidt/httprouter"
	"myfin/entity/web"
	"myfin/helper"
	"myfin/service"
	"net/http"
)

type PenarikanSaldoControllerImpl struct {
	PenarikanSaldoService service.PenarikanSaldoService
}

func NewPenarikanSaldoControllerImpl(penarikansaldoService service.PenarikanSaldoService) *PenarikanSaldoControllerImpl {
	return &PenarikanSaldoControllerImpl{
		PenarikanSaldoService: penarikansaldoService,
	}
}

func (controller *PenarikanSaldoControllerImpl) PenarikanSaldo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	penarikansaldoCreateRequest := web.PenarikanSaldoRequest{}
	helper.ReadFromRequestBody(request, &penarikansaldoCreateRequest)

	penarikansaldoResponse := controller.PenarikanSaldoService.PenarikanSaldo(request.Context(), penarikansaldoCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   penarikansaldoResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
