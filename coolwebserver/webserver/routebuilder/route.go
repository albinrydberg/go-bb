package routebuilder

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	route *mux.Route
}

func (r Route) WithPath(path string) Route {
	r.route.Path(path)
	return r
}

func (r Route) WithMethod(method string) Route {
	r.route.Methods(method)
	return r
}

func (r Route) WithHandler(handler func(http.ResponseWriter, *http.Request)) Route {
	r.route.HandlerFunc(handler)
	return r
}