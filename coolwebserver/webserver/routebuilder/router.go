package routebuilder

import (
	"github.com/gorilla/mux"
)

type Router struct {
	router *mux.Router
}

func NewRouter() Router {
	return Router{
		router: mux.NewRouter(),
	}
}

func (r Router) NewRoute() Route {
	return Route{
		route: r.router.NewRoute(),
	}
}

func (r Router) Router() *mux.Router {
	return r.router
}


