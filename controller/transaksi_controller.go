package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type TransaksiController interface {
	FindTransaksiKredit(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
