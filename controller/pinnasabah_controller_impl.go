package controller

import (
	"github.com/julienschmidt/httprouter"
	"myfin/entity/web"
	"myfin/helper"
	"myfin/service"
	"net/http"
	"strconv"
)

type PinNasabahControllerImpl struct {
	PinNasabahService service.PinNasabahService
}

func NewPinNasabahController(pinnasabahService service.PinNasabahService) *PinNasabahControllerImpl {
	return &PinNasabahControllerImpl{
		PinNasabahService: pinnasabahService,
	}
}

func (controller *PinNasabahControllerImpl) UpdatePin(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	datanasabahUpdatePinRequest := web.DatanasabahUpdatePinRequest{}
	helper.ReadFromRequestBody(request, &datanasabahUpdatePinRequest)

	datanasabahId := params.ByName("datanasabahId")
	id, err := strconv.Atoi(datanasabahId)
	helper.PanicIfError(err)

	datanasabahUpdatePinRequest.Id_user = id

	pinnasabahResponse := controller.PinNasabahService.UpdatePin(request.Context(), datanasabahUpdatePinRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   pinnasabahResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PinNasabahControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	datanasabahId := params.ByName("datanasabahId")
	id, err := strconv.Atoi(datanasabahId)
	helper.PanicIfError(err)

	pinnasabahResponse := controller.PinNasabahService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   pinnasabahResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
