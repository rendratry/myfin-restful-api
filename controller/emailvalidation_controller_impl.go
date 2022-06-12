package controller

import (
	"github.com/julienschmidt/httprouter"
	"myfin/entity/web"
	"myfin/helper"
	"myfin/service"
	"net/http"
	"strconv"
)

type EmailValidationControllerImpl struct {
	EmailValidationService service.EmailValidationService
}

func NewEmailValidationController(emailValidation service.EmailValidationService) *EmailValidationControllerImpl {
	return &EmailValidationControllerImpl{
		EmailValidationService: emailValidation,
	}
}

func (controller *EmailValidationControllerImpl) EmailValidation(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	emailValidationRequest := web.EmailValidationRequest{}
	helper.ReadFromRequestBody(request, &emailValidationRequest)

	emailValidationResponse := controller.EmailValidationService.EmailValidation(request.Context(), emailValidationRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   emailValidationResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *EmailValidationControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	emailValidationRequest := web.EmailValidationRequest{}
	helper.ReadFromRequestBody(request, &emailValidationRequest)

	emailvalidationId := params.ByName("emailvalidationId")
	id, err := strconv.Atoi(emailvalidationId)
	helper.PanicIfError(err)

	emailValidationRequest.IdUser = id

	emailvalidationResponse := controller.EmailValidationService.FindById(request.Context(), emailValidationRequest.IdUser)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   emailvalidationResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
