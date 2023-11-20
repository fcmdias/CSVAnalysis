package main

import (
	"context"
	"github/com/fcmdias/CSVAnalysis/services/backend/web"
	"log"
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

	server := &http.Server{
		Addr:    ":8080",
		Handler: wrappedMux,
	}

	go func() {
		log.Println("Starting server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	<-stopChan
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Server gracefully stopped")
}
