package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type AkunNasabahController interface {
	CreateAkun(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByEmail(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByNik(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
