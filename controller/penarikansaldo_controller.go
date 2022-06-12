package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PenarikanSaldoController interface {
	PenarikanSaldo(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
