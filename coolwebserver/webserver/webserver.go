package webserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"go-bb/coolwebserver/webserver/repohandler"
	"go-bb/coolwebserver/webserver/routebuilder"
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
	routeBuilder := routebuilder.NewRouter()

	routeBuilder.NewRoute().
		WithPath("/{key}").
		WithMethod(http.MethodGet).
		WithHandler(repoHandler.Get)

	routeBuilder.NewRoute().
		WithPath("/{key}").
		WithMethod(http.MethodPost).
		WithHandler(repoHandler.Post)

	return WebServer{
		router: routeBuilder.Router(),
	}
}

func (ws *WebServer) ListenAndServe(port int) {
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

func (ws *WebServer) Shutdown() {
	if err := ws.httpServer.Shutdown(context.Background()); err != nil {
		panic(err)
	}
}
