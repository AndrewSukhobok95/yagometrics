package database

import (
	"context"
	"database/sql"
	"log"

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
