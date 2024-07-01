package services

import (
	"context"
	"time"

	"go.opentelemetry.io/otel"
)

func SearchUser(ctx context.Context) {
	// Start a new child span
	_, span := otel.Tracer("service.find").Start(ctx, "SearchUser")
	defer span.End()

	time.Sleep(200 * time.Millisecond)
}
