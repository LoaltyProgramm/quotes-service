package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/LoaltyProgramm/quotes-service/internal/models/config"
)

func RunServer(cfg *config.Config) error {
	srv := http.Server{
		Addr: fmt.Sprintf(":%s", cfg.Port),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		}),
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	done := make(chan struct{})

	go func() {
		<-quit
		log.Println("signal received: shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("server forced to shutdown: %v", err)
		}

		log.Println("server exiting")

		close(done)
	}()

	log.Println("server is running on port", cfg.Port)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed{
		return fmt.Errorf("listen error: %w", err)
	}

	<-done
	return nil
}
