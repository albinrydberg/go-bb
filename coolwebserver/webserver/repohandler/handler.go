package repohandler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go-bb/coolwebserver/repository"
	"go-bb/coolwebserver/webserver/headers"
)

type RepoHandler struct {
	repo repository.Repository
}

func New(repo repository.Repository) RepoHandler {
	return RepoHandler{
		repo: repo,
	}
}

func (rh RepoHandler) Get(writer http.ResponseWriter, request *http.Request) {
	repoKey := mux.Vars(request)["key"]

	jpeg := rh.repo.Get(repoKey)

	writer.Header().Set(headers.ContentTypeHeaderKey, headers.ContentTypeJPEG)
	writer.Header().Set(headers.ContentLengthHeaderKey, fmt.Sprintf("%d", len(jpeg.Bytes())))

	if _, err := writer.Write(jpeg.Bytes()); err != nil {
		log.Println("unable to write image")
	}
}
