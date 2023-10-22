package database

import (
	"database/sql"
	"io/fs"

	"github.com/pressly/goose/v3"
)

const DefaultMigrationsFolder = "migrations"

func MigrateDB(db *sql.DB, dialect string, migrationsFS fs.FS) error {
	goose.SetBaseFS(migrationsFS)

	if err := goose.SetDialect(dialect); err != nil {
		return err
	}

	if err := goose.Up(db, DefaultMigrationsFolder); err != nil {
		return err
	}

	return nil
}
