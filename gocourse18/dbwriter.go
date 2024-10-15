package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBWriter struct {
	db *sqlx.DB
}

func initDBWriter(ctx context.Context, dbname string) (*DBWriter, error) {
	db, err := sqlx.ConnectContext(ctx, "postgres", "host=localhost port=5432 user=user password=password dbname=postgres sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("cannot connect to db: %v", err)
	}

	dbw := &DBWriter{
		db: db,
	}

	if dbname != "postgres" {
		err = dbw.createDatabaseIfNotExists(ctx, dbname)
		dbw.db.Close()
		if err != nil {
			return nil, err
		}

		dbw.db, err = sqlx.ConnectContext(ctx, "postgres", fmt.Sprintf("host=localhost port=5432 user=user password=password dbname=%s sslmode=disable", dbname))
		if err != nil {
			return nil, fmt.Errorf("cannot connect to db: %v", err)
		}
	}

	return dbw, nil
}

func (w *DBWriter) createDatabaseIfNotExists(ctx context.Context, dbName string) error {
	var existingDBName string
	err := w.db.QueryRowContext(ctx, "SELECT datname FROM pg_database WHERE datname = $1", dbName).Scan(&existingDBName)

	if err == sql.ErrNoRows {
		_, err = w.db.ExecContext(ctx, fmt.Sprintf("CREATE DATABASE %s", dbName))
		if err != nil {
			return fmt.Errorf("failed to create database: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("failed to check if database exists: %w", err)
	}

	return nil
}

func (w *DBWriter) createTables(ctx context.Context) error {
	_, err := w.db.ExecContext(ctx, schema)
	if err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}
	return nil
}

func (w *DBWriter) insertUsers(ctx context.Context, users []User) error {
	for _, user := range users {
		_, err := w.db.NamedExecContext(ctx, `INSERT INTO users (id, first_name, last_name, email, age, gender, city, trips_count, profession)
								VALUES (:id, :first_name, :last_name, :email, :age, :gender, :city, :trips_count, :profession)
								ON CONFLICT (email) DO NOTHING;`, &user)
		if err != nil {
			return fmt.Errorf("failed to insert user: %w", err)
		}
	}
	return nil
}

func (w *DBWriter) insertStatistics(ctx context.Context, statistics []Statistics) error {
	for _, stat := range statistics {
		_, err := w.db.NamedExecContext(ctx, `INSERT INTO statistics (city, age_range, average_trips)
								VALUES (:city, :age_range, :average_trips)
								ON CONFLICT (city, age_range) DO NOTHING;`, &stat)
		if err != nil {
			return fmt.Errorf("failed to insert statistics: %w", err)
		}
	}
	return nil
}

func (w *DBWriter) truncateTables(ctx context.Context) error {
	_, err := w.db.ExecContext(ctx, `TRUNCATE TABLE users, statistics RESTART IDENTITY;`)
	return err
}

func write(ctx context.Context, users []User, statistics []Statistics) error {
	dbw, err := initDBWriter(ctx, "taxi")
	if err != nil {
		return err
	}

	defer dbw.db.Close()

	err = dbw.createTables(ctx)
	if err != nil {
		return err
	}

	err = dbw.insertUsers(ctx, users)
	if err != nil {
		return err
	}

	err = dbw.insertStatistics(ctx, statistics)
	if err != nil {
		return err
	}

	return nil
}