package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type GetSaldoController interface {
	GetSaldo(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	MinSaldo(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
