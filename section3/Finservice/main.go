package main

import (
	"context"
	"errors"
	"github.com/go-chi/chi"
	"github.com/nade-harlow/go-test-exercise/section3/Finservice/handler"
	"github.com/nade-harlow/go-test-exercise/section3/Finservice/service/card"
	"github.com/nade-harlow/go-test-exercise/section3/Finservice/service/wallet"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	router := chi.NewMux()

	//chi logger middelware
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Request URI: %s", r.RequestURI)
			next.ServeHTTP(w, r)
		})
	})

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	wallet := wallet.NewWalletService()
	walletHandler := handler.NewWalletHandler(wallet)

	card := card.NewCardService()
	cardHandler := handler.NewCardHandler(card)

	// Group card routes
	router.Route("/api/v1/card", func(r chi.Router) {
		r.Post("/", cardHandler.CreateCard)
	})

	// Group wallet routes
	router.Route("/api/v1/wallet", func(r chi.Router) {
		r.Post("/", walletHandler.CreateWallet)
	})

	server := &http.Server{
		Addr:    ":8088",
		Handler: router,
	}

	go func() {
		log.Printf("Server listening on port 8088")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	// Graceful shutdown using signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop // Wait for SIGINT (Ctrl+C)

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	log.Printf("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Printf("Server stopped gracefully")
}
