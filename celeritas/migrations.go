package celeritas

import (
	"github.com/gobuffalo/pop"
	"log"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (c *Celeritas) popConnect() (*pop.Connection, error) {
	tx, err := pop.Connect("development")
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (c *Celeritas) CreatePopMigration(up, down []byte, migrationName, migrationType string) error {
	var migrationPath = c.RootPath + "/migrations"
	err := pop.MigrationCreate(migrationPath, migrationName, migrationType, up, down)
	if err != nil {
		return err
	}
	return nil
}

// Leaving the methods using golang-migrate here as a reference

func (c *Celeritas) MigrateUp(dsn string) error {
	m, err := migrate.New("file://"+c.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer func(m *migrate.Migrate) {
		err, _ := m.Close()
		if err != nil {

		}
	}(m)

	if err := m.Up(); err != nil {
		log.Println("Error running migration:", err)
		return err
	}
	return nil
}

func (c *Celeritas) MigrateDownAll(dsn string) error {
	m, err := migrate.New("file://"+c.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer func(m *migrate.Migrate) {
		err, _ := m.Close()
		if err != nil {

		}
	}(m)

	if err := m.Down(); err != nil {
		return err
	}

	return nil
}

func (c *Celeritas) Steps(n int, dsn string) error {
	m, err := migrate.New("file://"+c.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer func(m *migrate.Migrate) {
		err, _ := m.Close()
		if err != nil {

		}
	}(m)

	if err := m.Steps(n); err != nil {
		return err
	}

	return nil
}

func (c *Celeritas) MigrateForce(dsn string) error {
	m, err := migrate.New("file://"+c.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer func(m *migrate.Migrate) {
		err, _ := m.Close()
		if err != nil {

		}
	}(m)

	if err := m.Force(-1); err != nil {
		return err
	}

	return nil
}
