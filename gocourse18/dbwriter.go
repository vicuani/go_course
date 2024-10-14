package main

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func createDatabaseIfNotExists(db *sqlx.DB, dbName string) error {
	var existingDBName string
	err := db.QueryRow("SELECT datname FROM pg_database WHERE datname = $1", dbName).Scan(&existingDBName)

	if err == sql.ErrNoRows {
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
		if err != nil {
			return fmt.Errorf("failed to create database: %w", err)
		}
		fmt.Printf("DB '%s' created.\n", dbName)
	} else if err != nil {
		return fmt.Errorf("failed to check if database exists: %w", err)
	} else {
		fmt.Printf("DB '%s' already exists.\n", dbName)
	}

	return nil
}

func createTables(db *sqlx.DB) error {
	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}
	return nil
}

func insertUsers(db *sqlx.DB, users []User) error {
	for _, user := range users {
		_, err := db.NamedExec(`INSERT INTO users (first_name, last_name, email, age, gender, city, trips_count, profession)
								VALUES (:first_name, :last_name, :email, :age, :gender, :city, :trips_count, :profession)
								ON CONFLICT (email) DO NOTHING;`, &user)
		if err != nil {
			return fmt.Errorf("failed to insert user: %w", err)
		}
	}
	fmt.Println("Users added successfully")
	return nil
}

func insertStatistics(db *sqlx.DB, statistics []Statistics) error {
	for _, stat := range statistics {
		_, err := db.NamedExec(`INSERT INTO statistics (city, age_range, average_trips)
								VALUES (:city, :age_range, :average_trips)
								ON CONFLICT (city, age_range) DO NOTHING;`, &stat)
		if err != nil {
			return fmt.Errorf("failed to insert statistics: %w", err)
		}
	}
	fmt.Println("Statistics added successfully")
	return nil
}

func truncateTables(db *sqlx.DB) error {
	_, err := db.Exec(`TRUNCATE TABLE users, statistics RESTART IDENTITY;`)
	return err
}

func write(users []User, statistics []Statistics) error {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=user password=password dbname=postgres sslmode=disable")
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}
	defer db.Close()

	err = createDatabaseIfNotExists(db, "taxi")
	if err != nil {
		return err
	}

	db, err = sqlx.Connect("postgres", "host=localhost port=5432 user=user password=password dbname=taxi sslmode=disable")
	if err != nil {
		return fmt.Errorf("failed to connect to taxi database: %w", err)
	}
	defer db.Close()

	err = createTables(db)
	if err != nil {
		return err
	}

	err = insertUsers(db, users)
	if err != nil {
		return err
	}

	err = insertStatistics(db, statistics)
	if err != nil {
		return err
	}

	return nil
}
