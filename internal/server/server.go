package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"project8/internal/hndlr"
	"syscall"
	"time"
)

func Start(address string) {

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/image-download-size", hndlr.BaseHandler)
	mux.HandleFunc("/", hndlr.NotFoundUrl)

	server := http.Server{
		Addr:    address,
		Handler: mux,
	}

	log.Printf("Сервер запущен. Адрес: %s. PID: %d\n", address, os.Getppid())
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("Error: %v\n", err)
		}
	}()

	<-ctx.Done()
	shutdownCtx, _ := context.WithTimeout(context.Background(), 7*time.Second)
	server.Shutdown(shutdownCtx)
	log.Println("Сервер закрыт.")
}
