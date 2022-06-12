package controller

import (
	"github.com/julienschmidt/httprouter"
	"myfin/entity/web"
	"myfin/helper"
	"myfin/service"
	"net/http"
)

type LoginNasabahControllerImpl struct {
	LoginNasabahService service.LoginNasabahService
}

func NewLoginNasabahController(loginnasabahService service.LoginNasabahService) *LoginNasabahControllerImpl {
	return &LoginNasabahControllerImpl{
		LoginNasabahService: loginnasabahService,
	}
}

func (controller *LoginNasabahControllerImpl) FindByEmail(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginnasabahRequest := web.LoginNasabahRequest{}
	helper.ReadFromRequestBody(request, &loginnasabahRequest)

	loginnasabahResponse := controller.LoginNasabahService.FindByEmail(request.Context(), loginnasabahRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   loginnasabahResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
