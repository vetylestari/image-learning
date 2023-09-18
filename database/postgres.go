package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/stdlib"
)

func Open() *sql.DB {
	//Init DB
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_SCHEMA"),
	)
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func CreateSchema() (sql.Result, error) {
	//Init DB
	conn := Open()
	defer conn.Close()
	return conn.ExecContext(context.Background(), fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", os.Getenv("DB_SCHEMA")))
}
