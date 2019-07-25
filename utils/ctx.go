package utils

import (
	"context"
	"time"
)

// GetCtx .
func GetCtx() context.Context {
	ctxWithTimeout, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ctxWithTimeout
}
