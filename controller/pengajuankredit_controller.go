package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PengajuanKreditController interface {
	AjuanKredit(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
