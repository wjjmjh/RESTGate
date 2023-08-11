package service

import (
	"RESTGate/pkg/models"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func NewRouter(services []models.Service) *mux.Router {

	r := mux.NewRouter().StrictSlash(true)

	for _, service := range services {
		for _, route := range service.Routes {
			r.Methods(strings.Split(route.Method, ",")...).
				Path(route.Pattern).
				Name(route.Name).
				Handler(http.HandlerFunc(ReverseHandlerFactory(service.ServiceUrl)))
		}
	}

	return r
}
