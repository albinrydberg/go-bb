package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"go-bb/coolwebserver/repository"
	"go-bb/coolwebserver/webserver"
	"go-bb/coolwebserver/webserver/repohandler"
	"go-bb/imageloader/smartimageloader/catloader"
)

const (
	serverPort = 8080
)

func main() {
	repoHandler := repohandler.New(
		repository.NewPreLoaded("cat", catloader.CatLoader{}),
	)

	server := webserver.New(repoHandler)

	signals := make(chan os.Signal, 1)
	serverClosed := make(chan struct{}, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signals
		log.Println("Will try to exit gracefully")
		server.Shutdown()
		close(serverClosed)
	}()

	go server.ListenAndServe(serverPort)

	<-serverClosed
	log.Println("Exiting")
}
