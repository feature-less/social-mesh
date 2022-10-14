// spawn services with features such as:
// graceful shutdown, and restart
package service

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"server/config"
	"server/internal/probe"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

// BaseContext will always be overridden
type Service config.ConfigureService

// spawns a new service with configurations provided through
// service struct
func (service *Service) Spawn() {
	// make sure we print the pid of this server
	log.Printf("Process ID: %d", os.Getpid())

	// meant so that when errgroup receives an error
	// which mostly is meant for graceful shutdown of the service
	// the children would be notified
	// so that they do their proper measures before fully shutting down
	ctx, stop := signal.NotifyContext(context.Background())

	// if for whatever reasons the server starts
	// but the function goes through it as if it didn't
	// and it doesn't return any errors so the service is never shutdown properly
	// we make sure this safeguard executes
	defer stop()

	// needed to notify children of service shuttingdown
	service.HTTP.BaseContext = func(_ net.Listener) context.Context {
		return ctx
	}

	// create new http server config for our service
	var server http.Server = service.HTTP

	// create a new livefile config for our service
	var liveFile probe.LiveFile

	// set path to tmp
	// not to be confused with the regular tmp directory
	// this file will be in our current working diractory
	liveFile.Path = "tmp"

	// the basename of our temp file should always be unique
	// so that we could have the same worker spawn multiple srvices
	// ex: Accounts Service(this is the name of our worker) is running on ports: 3000, 3001, 3002
	// eash of these services will have a completely unique livefile
	liveFile.BaseName = uuid.NewString()

	// make sure to remove the livefile only when the function is returning
	// not when the server is shut odwn
	// because livefiles are meant to check if a process is stil running or not
	// for http health check create an endpoint ex: /health
	// and use that to check if the http server is still up
	defer func() {
		// check if livefile exists first
		liveFileExists := liveFile.Exists()

		// we attempt to remove the livefile
		if liveFileExists == true {
			if err := liveFile.Remove(); err != nil {
				log.Fatalf("an unexpected error has occured while removing live file: %v", err)
			}
			return
		}
		log.Fatalf("Fatal: could not find livefile associated with this service: \n livefile path: %v \n http address associated with the service: %v", liveFile.Path+"/"+liveFile.BaseName, server.Addr)
		return
	}()

	// handle signals that meant to shutdown or restart the service
	go func() {
		// create a channel for incoming signals
		sigint := make(chan os.Signal, 1)

		// notify listeners of signal interupt
		// mostly meant for direct manual administration of services
		// such as during development
		signal.Notify(sigint, os.Interrupt)

		// notify listners of SIGNTERM
		// mostly meant for k8s and such technologies
		signal.Notify(sigint, syscall.SIGTERM)

		<-sigint
		stop()

	}()

	// this is an errgroup using the same context that we have passed to our
	// httpserver in its BaseContext
	errorGroup, errorGroupContext := errgroup.WithContext(ctx)

	// goroutine to start/create anything related to service startup
	errorGroup.Go(func() error {
		// attempt to create this service's livefile
		_, err := liveFile.Create()
		if err != nil {
			log.Fatalf("an unexpected Error has occured while creating live file: %v", err)
		}

		// log that our srvice is running
		log.Printf("Staring service at: %v", service.HTTP.Addr)

		// this function will create an infinite loop
		// and lock our function until the service shuts down
		// therefore if you want anything to run before this function
		// put it before it anything after won't get executed until shutdown
		err = server.ListenAndServe()

		// as was said before the server will lock and not return
		// until it has fully shutdown
		if err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	// gorouting to cleanup after our service and shut it down
	errorGroup.Go(func() error {
		<-errorGroupContext.Done()

		// log that our service is shutting down
		log.Printf("Service shutting down at: %v", service.HTTP.Addr)
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatalf("an unexpected error has occured while shutting down server: %v", err)
		}
		return nil
	})

	// the errgroup has no business to keep waiting when all other goroutines are done
	// so we log the error and exit
	if err := errorGroup.Wait(); err != nil {
		log.Fatalf("an unexpected error has occured: %v", err)
	}
}
