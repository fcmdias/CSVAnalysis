package main

import (
	"context"
	"github/com/fcmdias/CSVAnalysis/services/backend/web"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	router := http.NewServeMux()
	router.Handle("/popular", web.LoggingMiddleware(web.EnableCORSMiddleware(http.HandlerFunc(web.PopularHandler))))
	router.Handle("/byyear", web.LoggingMiddleware(web.EnableCORSMiddleware(http.HandlerFunc(web.ByYearHandler))))
	router.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	wrappedMux := web.PanicRecoveryMiddleware(router)

	port := "8080"
	server := &http.Server{
		Addr:    ":" + port,
		Handler: wrappedMux,
	}

	go func() {
		slog.Info("Starting server", "port", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("ListenAndServe", "error", err)
		}
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	<-stopChan
	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server Shutdown Failed", "error", err)
	}
	slog.Info("Server gracefully stopped")
}
