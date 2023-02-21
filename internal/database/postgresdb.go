package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/datastorage"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/serialization"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/mattn/go-sqlite3"
)

type PostgresDB struct {
	db *sql.DB
}

func NewDB(address string) *PostgresDB {
	db, err := sql.Open("pgx", address)
	//db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		log.Printf("Couldn't connet to DB:\n")
		panic(err)
	}
	return &PostgresDB{db: db}
}

func (pgdb *PostgresDB) Close() {
	pgdb.db.Close()
}

func (pgdb *PostgresDB) PingContext(ctx context.Context) error {
	err := pgdb.db.PingContext(ctx)
	return err
}

func (pgdb *PostgresDB) CreateTable(ctx context.Context) {
	query := `
		CREATE TABLE IF NOT EXISTS public.metrics (
			id text not null, 
			type text not null,
			value double precision,
			delta BIGINT
		);
	`
	_, err := pgdb.db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error while creting table:\n")
		log.Printf(err.Error() + "\n")
		return
	}
}

func (pgdb *PostgresDB) UpdateMetricInDB(metric serialization.Metrics, ctx context.Context) error {
	q := `
	UPDATE metrics SET
	id=$1, type=$2, value=$3, delta=$4  
	WHERE id=$1 and type=$2;
	`
	commandTag, _ := pgdb.db.ExecContext(ctx, q, metric.ID, metric.MType, metric.Value, metric.Delta)
	rows, err := commandTag.RowsAffected()
	if err != nil {
		log.Printf(err.Error() + "\n\n")
	}

	if rows != 1 {
		q = `
		INSERT INTO metrics (id, type, value, delta)
		SELECT $1, $2, $3, $4 
		WHERE NOT EXISTS (SELECT 1 FROM metrics WHERE id=$1);
		`
		_, err = pgdb.db.ExecContext(ctx, q, metric.ID, metric.MType, metric.Value, metric.Delta)
		if err != nil {
			log.Printf(err.Error() + "\n\n")
		}
	}
	return nil
}

func (pgdb *PostgresDB) UpdateDB(storage datastorage.Storage, storeInterval time.Duration, ctx context.Context) {
	ticker := time.NewTicker(storeInterval)
	for {
		<-ticker.C

		data := storage.ExportToJSON()
		var metrics []serialization.Metrics
		err := json.Unmarshal(data, &metrics)
		if err != nil {
			log.Printf(err.Error() + "\n\n")
			return
		}

		for _, m := range metrics {
			err = pgdb.UpdateMetricInDB(m, ctx)
			if err != nil {
				log.Printf(err.Error() + "\n\n")
				return
			}
		}
		log.Println("DB is successfully updated")
	}
}

func (pgdb *PostgresDB) StartWritingToDB(storage datastorage.Storage, storeInterval time.Duration, ctx context.Context, wg *sync.WaitGroup) {
	pgdb.CreateTable(ctx)

	log.Printf("Updating DB every " + storeInterval.String() + "\n")
	wg.Add(1)
	go func() {
		pgdb.UpdateDB(storage, storeInterval, ctx)
	}()
}
