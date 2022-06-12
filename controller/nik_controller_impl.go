package controller

import (
	"github.com/julienschmidt/httprouter"
	"myfin/entity/web"
	"myfin/helper"
	"myfin/service"
	"net/http"
)

type GetNikControllerImpl struct {
	GetNikService service.GetNikService
}

func NewGetNikController(getNik service.GetNikService) *GetNikControllerImpl {
	return &GetNikControllerImpl{
		GetNikService: getNik,
	}
}

func (controller *GetNikControllerImpl) GetNik(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	getNikRequest := web.NikRequest{}
	helper.ReadFromRequestBody(request, &getNikRequest)

	getNikResponse := controller.GetNikService.GetNik(request.Context(), getNikRequest.Nik)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   getNikResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
