package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type GetProfileController interface {
	GetProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
