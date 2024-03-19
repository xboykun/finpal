package main

import (
	"errors"
	"log"
	"nami/internal/delivery/http/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(
		stopCh,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	var (
		addr = "localhost:5678"
		mux  = http.NewServeMux()
		srv  = http.Server{
			Addr:              addr,
			Handler:           mux,
			ReadTimeout:       5 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			WriteTimeout:      5 * time.Second,
		}
	)

	router.AccountRouter(mux)

	go func() {
		log.Printf("start http server, addr: http://%s\n", addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("stop http server, err: %v\n", err)
		}
	}()

	<-stopCh
	log.Printf("stop http server, addr: http://%s\n", addr)
}
