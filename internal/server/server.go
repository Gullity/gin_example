package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	v1 "github.com/Gullity/gin_example/internal/api/v1"
	"github.com/Gullity/gin_example/internal/config"
	"github.com/Gullity/gin_example/internal/otel"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func StartServer(ctx context.Context, tracerShutdown, metricShutdown otel.OtelFunc) {
	engine := gin.Default()

	engine.Use(otelgin.Middleware(config.Config.Service.Name))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Config.Service.Port),
		Handler: engine.Handler(),
	}

	router := engine.Group(v1.ApiVersion)

	usersRouters := v1.NewUsers()
	router.GET(usersRouters.Path(), usersRouters.Find)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	go signalsListener(srv, ctx, tracerShutdown, metricShutdown)

	// Block the main goroutine to keep the server running
	select {}
}

func signalsListener(server *http.Server, ctx context.Context, traceShutdown, metricShutdown otel.OtelFunc) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	sig := <-sigs

	log.Printf("Received signal: %s. Initiating graceful shutdown...", sig)

	gracefulStop(server, ctx, traceShutdown, metricShutdown)
}

func gracefulStop(server *http.Server, ctx context.Context, traceShutdown, metricShutdown otel.OtelFunc) {
	ctxShutdown, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	traceShutdown(ctxShutdown)
	metricShutdown(ctxShutdown)

	log.Println("Server exiting")
	os.Exit(0)
}
