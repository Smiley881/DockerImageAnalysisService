package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"project8/internal/hndlr"
	"syscall"
	"time"
)

func Start(address string) {
	var shutdownTime time.Duration
	var err error

	shutdownTimeStr, exists := os.LookupEnv("SHUTDOWN_TIME")
	if exists {
		shutdownTime, err = time.ParseDuration(shutdownTimeStr)
		if err != nil {
			log.Panicf("Error: Ошибка валидации переменной окружения SHUTDOWN_TIME: %v", err)
		}
	} else {
		shutdownTime = 7 * time.Second
	}

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
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Error: %v\n", err)
		}
	}()

	<-ctx.Done()
	shutdownCtx, _ := context.WithTimeout(context.Background(), shutdownTime)
	server.Shutdown(shutdownCtx)
	log.Println("Сервер закрыт.")
}
