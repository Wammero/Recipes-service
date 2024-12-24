package migrate

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func CallMigrations(dsn string) (*migrate.Migrate, error) {
	migrationsPath := "file://migrations"

	m, err := migrate.New(migrationsPath, dsn)

	if err != nil {
		return nil, err
	}

	return m, err
}
