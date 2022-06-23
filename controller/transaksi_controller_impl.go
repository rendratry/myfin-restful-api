package controller

import (
	"github.com/julienschmidt/httprouter"
	"myfin/entity/web"
	"myfin/helper"
	"myfin/service"
	"net/http"
	"strconv"
)

type TransaksiControllerImpl struct {
	TransaksiService service.TransaksiService
}

func NewTransaksiController(transaksiService service.TransaksiService) *TransaksiControllerImpl {
	return &TransaksiControllerImpl{
		TransaksiService: transaksiService,
	}
}

func (controller *TransaksiControllerImpl) FindTransaksiKredit(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transaksiRequest := web.TransaksiRequest{}

	nasabahId := params.ByName("nasabahId")
	status := params.ByName("status")
	id, err := strconv.Atoi(nasabahId)
	helper.PanicIfError(err)

	transaksiRequest.IdNasabah = id
	transaksiRequest.Status = status

	avaResponse := controller.TransaksiService.FindTransaksiKredit(request.Context(), transaksiRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   avaResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
