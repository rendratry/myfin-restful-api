package controller

import (
	"github.com/julienschmidt/httprouter"
	"myfin/entity/web"
	"myfin/helper"
	"myfin/service"
	"net/http"
	"strconv"
)

type AvaUploadControllerImpl struct {
	AvaUploadService service.AvaUploadService
}

func NewAvaUploadController(avauploadService service.AvaUploadService) *AvaUploadControllerImpl {
	return &AvaUploadControllerImpl{
		AvaUploadService: avauploadService,
	}
}

func (controller *AvaUploadControllerImpl) AvaUpload(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	AvaUploadRequest := web.AvaUploadRequest{}
	helper.ReadFromRequestBody(request, &AvaUploadRequest)

	nasabahId := params.ByName("nasabahId")
	id, err := strconv.Atoi(nasabahId)
	helper.PanicIfError(err)

	AvaUploadRequest.Id_user = id

	avauploadResponse := controller.AvaUploadService.AvaUpload(request.Context(), AvaUploadRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   avauploadResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AvaUploadControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	datanasabahId := params.ByName("nasabahId")
	id, err := strconv.Atoi(datanasabahId)
	helper.PanicIfError(err)

	avaResponse := controller.AvaUploadService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   avaResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
