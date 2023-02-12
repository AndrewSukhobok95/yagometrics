package database

import (
	"context"
)

type CustomDB interface {
	Close()
	PingContext(ctx context.Context) error
}
