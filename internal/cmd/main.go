package main

import (
	"context"
	"log"

	"github.com/Gullity/gin_example/internal/config"
	"github.com/Gullity/gin_example/internal/otel"
	"github.com/Gullity/gin_example/internal/server"
)

func main() {
	ctx := context.Background()
	config.InitConfig()

	traceShutdown, metricShutdown, err := otel.StartOtel(ctx)
	if err != nil {
		log.Fatalf("otel failed error: %s", err)
	}

	server.StartServer(ctx, traceShutdown, metricShutdown)

}
