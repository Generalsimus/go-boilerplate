package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Generalsimus/go-monolith-boilerplate/config"
	"github.com/Generalsimus/go-monolith-boilerplate/db/database"
	"github.com/Generalsimus/go-monolith-boilerplate/internal/handlers"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	pool, err := pgxpool.New(context.Background(), config.Cfg.DATABASE_URL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer pool.Close()

	// 2. Instantiate the HTTP Server explicitly
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Cfg.PORT),
		Handler: handlers.Routes(database.New(pool)),
	}
	// 3. Start the server in a separate goroutine
	go func() {
		log.Println(fmt.Sprintf("🚀 Server is running on :%d", config.Cfg.PORT))
		// ErrServerClosed is expected when we call srv.Shutdown()
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server crashed: %v", err)
		}
	}()
	// 4. Create a channel to listen for OS signals
	quit := make(chan os.Signal, 1)

	// Catch SIGINT (Ctrl+C) and SIGTERM (Docker/Kubernetes shutdown)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 5. Block the main thread here until a signal is received
	<-quit
	log.Println("🛑 Shutting down server...")

	// 6. Create a context with a timeout to force shutdown if it takes too long
	// 10 seconds is usually enough time for ongoing HTTP requests to finish
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// 7. Execute the graceful shutdown
	// This stops accepting new connections and waits for active ones to finish
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("✅ Server exited cleanly")
}
