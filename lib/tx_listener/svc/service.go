package svc

import (
	"context"
)

// Service is an interface that defines the methods for a transactions listener service.
type Service interface {
	// StartListener is a method that starts the listener service and returns an error if any occurs.
	StartListener(ctx context.Context) error
}
