package controller

import (
	"github.com/julienschmidt/httprouter"
	"myfin/entity/web"
	"myfin/helper"
	"myfin/service"
	"net/http"
	"strconv"
)

type KeamananNasabahControllerImpl struct {
	KeamananNasabahService service.KeamananNasabahService
}

func NewKeamananNasabahController(keamanannasabahService service.KeamananNasabahService) *KeamananNasabahControllerImpl {
	return &KeamananNasabahControllerImpl{
		KeamananNasabahService: keamanannasabahService,
	}
}

func (controller *KeamananNasabahControllerImpl) UpdateKeamanan(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	datanasabahUpdateKeamananRequest := web.DatanasabahUpdateKeamananRequest{}
	helper.ReadFromRequestBody(request, &datanasabahUpdateKeamananRequest)

	datanasabahId := params.ByName("datanasabahId")
	id, err := strconv.Atoi(datanasabahId)
	helper.PanicIfError(err)

	datanasabahUpdateKeamananRequest.Id_user = id

	datanasabahResponse := controller.KeamananNasabahService.UpdateKeamanan(request.Context(), datanasabahUpdateKeamananRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   datanasabahResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *KeamananNasabahControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	datanasabahCekKeamananRequest := web.DatanasabahUpdateKeamananRequest{}
	helper.ReadFromRequestBody(request, &datanasabahCekKeamananRequest)

	datanasabahId := params.ByName("datanasabahId")
	id, err := strconv.Atoi(datanasabahId)
	helper.PanicIfError(err)

	datanasabahCekKeamananRequest.Id_user = id

	keamananResponse := controller.KeamananNasabahService.FindScurity(request.Context(), datanasabahCekKeamananRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   keamananResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
