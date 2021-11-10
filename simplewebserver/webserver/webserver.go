package webserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"go-bb/simplewebserver/webserver/simpleimagehandler"
)

const (
	readTimeout  = 30 * time.Second
	writeTimeout = 30 * time.Second
)

type WebServer struct {
	httpServer *http.Server
	router     *mux.Router
}

func New() *WebServer {
	router := mux.NewRouter()

	route := router.NewRoute()
	route.Path("/")
	route.Methods(http.MethodGet)
	route.HandlerFunc(simpleimagehandler.Handle)

	return &WebServer{
		router: router,
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
