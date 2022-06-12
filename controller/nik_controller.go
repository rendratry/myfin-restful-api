package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type GetNikController interface {
	GetNik(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
