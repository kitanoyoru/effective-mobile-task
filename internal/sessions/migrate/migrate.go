package migrate

import (
	"database/sql"
	"embed"

	"github.com/kitanoyoru/effective-mobile-task/internal/config"
	"github.com/kitanoyoru/effective-mobile-task/pkg/database"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

type MigrateSession struct {
	providerName string
	db           *sql.DB
}

func NewMigrateSession(cfg *config.DatabaseConfig) (*MigrateSession, error) {
	gdb, err := database.ConnectToDB(cfg)
	if err != nil {
		return nil, err
	}

	db, err := gdb.DB()
	if err != nil {
		return nil, err
	}

	return &MigrateSession{
		cfg.Name,
		db,
	}, nil
}

func (m *MigrateSession) Migrate() error {
	return database.MigrateDB(m.db, m.providerName, embedMigrations)
}

func (m *MigrateSession) Close() error {
	return m.db.Close()
}
