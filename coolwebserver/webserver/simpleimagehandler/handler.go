package simpleimagehandler

import (
	"fmt"
	"log"
	"net/http"

	"go-bb/coolwebserver/webserver/headers"
	"go-bb/imageloader"
)

func Handle(writer http.ResponseWriter, _ *http.Request) {
	jpeg, err := imageloader.LoadGoat()
	if err != nil {
		panic(err)
	}

	writer.Header().Set(headers.ContentTypeHeaderKey, headers.ContentTypeJPEG)
	writer.Header().Set(headers.ContentLengthHeaderKey, fmt.Sprintf("%d", len(jpeg.Bytes())))

	if _, err := writer.Write(jpeg.Bytes()); err != nil {
		log.Println("unable to write image")
	}
}
