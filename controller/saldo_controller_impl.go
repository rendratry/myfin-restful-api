package controller

import (
	"github.com/julienschmidt/httprouter"
	"myfin/entity/web"
	"myfin/helper"
	"myfin/service"
	"net/http"
	"strconv"
)

type GetSaldoControllerImpl struct {
	GetSaldoService service.GetSaldoService
}

func NewGetSaldoController(getsaldoService service.GetSaldoService) *GetSaldoControllerImpl {
	return &GetSaldoControllerImpl{
		GetSaldoService: getsaldoService,
	}
}

func (controller *GetSaldoControllerImpl) GetSaldo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	saldoId := params.ByName("saldoId")
	id, err := strconv.Atoi(saldoId)
	helper.PanicIfError(err)

	saldoResponse := controller.GetSaldoService.GetSaldo(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   saldoResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *GetSaldoControllerImpl) MinSaldo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	minsaldoUpdateRequest := web.MinSaldoUpdateRequest{}
	helper.ReadFromRequestBody(request, &minsaldoUpdateRequest)

	saldoId := params.ByName("saldoId")
	id, err := strconv.Atoi(saldoId)
	helper.PanicIfError(err)

	minsaldoUpdateRequest.IdNasabah = id

	minsaldoResponse := controller.GetSaldoService.MinSaldo(request.Context(), minsaldoUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   minsaldoResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
