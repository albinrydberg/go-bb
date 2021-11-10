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
	"go-bb/imageloader/smartimageloader/dogloader"
	"go-bb/imageloader/smartimageloader/goatloader"
	"go-bb/imageloader/smartimageloader/gopherloader"
)

const (
	serverPort = 8080
)

func main() {
	server := webserver.New(
		repohandler.New(
			initRepository(),
		),
	)

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

func initRepository() repository.Repository {
	repo := repository.NewWithDefault(gopherloader.GopherLoader{})

	if err := repo.Load("cat", catloader.CatLoader{}); err != nil {
		panic(err)
	}

	if err := repo.Load("dog", dogloader.DogLoader{}); err != nil {
		panic(err)
	}

	if err := repo.Load("goat", goatloader.GoatLoader{}); err != nil {
		panic(err)
	}

	return repo
}
