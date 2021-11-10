package webserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"go-bb/coolwebserver/webserver/repohandler"
	"go-bb/coolwebserver/webserver/simpleimagehandler"
)

const (
	readTimeout  = 30 * time.Second
	writeTimeout = 30 * time.Second
)

type WebServer struct {
	httpServer *http.Server
	router     *mux.Router
}

func New(repoHandler repohandler.RepoHandler) WebServer {
	router := mux.NewRouter()

	route := router.NewRoute()
	route.Path("/")
	route.Methods(http.MethodGet)
	route.HandlerFunc(simpleimagehandler.Handle)

	route2 := router.NewRoute()
	route2.Path("/{key}")
	route2.Methods(http.MethodGet)
	route2.HandlerFunc(repoHandler.Get)

	return WebServer{
		router: router,
	}
}

func (ws WebServer) ListenAndServe(port int) {
	ws.httpServer = &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      ws.router,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	if err := ws.httpServer.ListenAndServe(); err != http.ErrServerClosed {
		panic(err)
	}
}

func (ws WebServer) Shutdown() {
	if err := ws.httpServer.Shutdown(context.Background()); err != nil {
		panic(err)
	}
}
