package controller

import (
	"github.com/julienschmidt/httprouter"
	"myfin/entity/web"
	"myfin/helper"
	"myfin/service"
	"net/http"
)

type AkunNasabahControllerImpl struct {
	AkunNasabahService service.AkunNasabahService
}

func NewAkunNasabahController(akunnasabahService service.AkunNasabahService) *AkunNasabahControllerImpl {
	return &AkunNasabahControllerImpl{
		AkunNasabahService: akunnasabahService,
	}
}

func (controller *AkunNasabahControllerImpl) CreateAkun(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	akunnasabahCreateRequest := web.AkunNasabahCreateRequest{}
	helper.ReadFromRequestBody(request, &akunnasabahCreateRequest)

	akunnasabahResponse := controller.AkunNasabahService.CreateAkun(request.Context(), akunnasabahCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   akunnasabahResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
