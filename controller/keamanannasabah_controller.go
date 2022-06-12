package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type KeamananNasabahController interface {
	UpdateKeamanan(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
