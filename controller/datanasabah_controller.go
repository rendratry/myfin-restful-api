package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type DatanasabahController interface {
	UpdateDataNasabah(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
