package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"post/internal/config"

	"post/internal"
)

func main() {
	opt := config.Options{
		ListenAddress:   os.Getenv("LISTEN_ADDRESS"),
		DatabaseDSN:     os.Getenv("DATABASE_DSN"),
		CORSAllowOrigin: strings.Split(os.Getenv("CORS_ALLOW_ORIGIN"), ","),
	}

	s, err := internal.New(opt)

	if err != nil {
		log.Fatal("internal.New():", err)
	}

	srv := &http.Server{
		Addr:    opt.ListenAddress,
		Handler: s.Router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Gracefully shutdown
	gracefullyShutdownOnSignal(srv, s)

	srv.Close()
}

func gracefullyShutdownOnSignal(server *http.Server, service *internal.Service) {
	sig := waitForShutdownSignal()
	log.Println(sig + ": Received signal, starting shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Error shutting down server:", err)
	} else {
		log.Println("Server gracefully stopped")
	}

	db, _ := service.ConnDB.DB()

	if err := db.Close(); err != nil {
		log.Fatal("Error clost db connection:", err)
	} else {
		log.Println("DB connection gracefully closed")
	}
}

func waitForShutdownSignal() string {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	sig := <-quit

	return sig.String()
}
