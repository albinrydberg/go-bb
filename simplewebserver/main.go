package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"go-bb/simplewebserver/webserver"
)

const (
	serverPort = 8080
)

func main() {
	server := webserver.New()

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
