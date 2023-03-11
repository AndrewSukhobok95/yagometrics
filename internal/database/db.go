package database

import (
	"context"
	"sync"
	"time"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/datastorage"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/serialization"
)

type CustomDB interface {
	Close()
	PingContext(ctx context.Context) error
	CreateTable(ctx context.Context)
	UpdateMetricInDB(metric serialization.Metrics, ctx context.Context) error
	UpdateDB(storage datastorage.Storage, storeInterval time.Duration, ctx context.Context)
	StartWritingToDB(storage datastorage.Storage, storeInterval time.Duration, ctx context.Context, wg *sync.WaitGroup)
}
