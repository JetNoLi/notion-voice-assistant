package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jetnoli/notion-voice-assistant/config"
)

var (
	once sync.Once
	Db   *pgxpool.Pool
)

type QueryFunc = func(string) (pgx.Rows, error)

func Connect() QueryFunc {
	once.Do(func() {
		// Connection close happens in main
		conn, err := pgxpool.New(context.Background(), config.DBConnectionString)
		if err != nil {
			panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
		}

		Db = conn
	})

	return Query
}

func Query(queryString string) (pgx.Rows, error) {
	return Db.Query(context.Background(), queryString)
}

func QueryRow(queryString string) pgx.Row {
	return Db.QueryRow(context.Background(), queryString)
}
