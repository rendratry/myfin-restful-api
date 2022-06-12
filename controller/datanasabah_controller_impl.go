package controller

import (
	"github.com/julienschmidt/httprouter"
	"myfin/entity/web"
	"myfin/helper"
	"myfin/service"
	"net/http"
	"strconv"
)

type DatanasabahControllerImpl struct {
	DatanasabahService service.DataNasabahService
}

func NewDatanasabahController(datanasabahService service.DataNasabahService) *DatanasabahControllerImpl {
	return &DatanasabahControllerImpl{
		DatanasabahService: datanasabahService,
	}
}

func (controller *DatanasabahControllerImpl) UpdateDataNasabah(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	datanasabahUpdateRequest := web.DatanasabahUpdateRequest{}
	helper.ReadFromRequestBody(request, &datanasabahUpdateRequest)

	datanasabahId := params.ByName("datanasabahId")
	id, err := strconv.Atoi(datanasabahId)
	helper.PanicIfError(err)

	datanasabahUpdateRequest.Id_user = id

	datanasabahResponse := controller.DatanasabahService.UpdateDataNasabah(request.Context(), datanasabahUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   datanasabahResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DatanasabahControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	datanasabahId := params.ByName("id_user")
	id, err := strconv.Atoi(datanasabahId)
	helper.PanicIfError(err)

	categoryResponse := controller.DatanasabahService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
