package database

import (
	"context"
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

func Migration(ctx context.Context, connStr string) {
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("Failed to create driver:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://board/internal/migrations",
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatal("Failed to create migrate instance:", err)
	}

	// Проверяем статус миграций
	version, dirty, err := m.Version()
	if err != nil {
		log.Printf("Cannot get migration version: %v", err)
	} else {
		log.Printf("Current version: %d, dirty: %t", version, dirty)
	}

	// Если база "грязная", фиксируем версию
	if dirty {
		log.Printf("Fixing dirty database version %d", version)
		err = m.Force(int(version))
		if err != nil {
			log.Fatal("Failed to force version:", err)
		}
		log.Println("Dirty version fixed, retrying migration...")
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migrations applied successfully")

}
