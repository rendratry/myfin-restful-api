package controller

import (
	"github.com/julienschmidt/httprouter"
	"myfin/entity/web"
	"myfin/helper"
	"myfin/service"
	"net/http"
	"strconv"
)

type GetProfileControllerImpl struct {
	GetProfileService service.GetProfileService
}

func NewGetProfileController(getProfile service.GetProfileService) *GetProfileControllerImpl {
	return &GetProfileControllerImpl{
		GetProfileService: getProfile,
	}
}

func (controller *GetProfileControllerImpl) GetProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	datanasabahId := params.ByName("datanasabahId")
	id, err := strconv.Atoi(datanasabahId)
	helper.PanicIfError(err)

	pinnasabahResponse := controller.GetProfileService.GetProfile(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   pinnasabahResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
