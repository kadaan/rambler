package firebirdsql

import (
	"database/sql"
	"fmt"

	"github.com/elwinar/rambler/driver"
	"github.com/elwinar/rambler/env"
	_ "github.com/nakagami/firebirdsql"
)

func init() {
	driver.Register("firebirdsql", Driver{})
}

// Driver is the type that initialize new connections.
type Driver struct{}

func (d Driver) New(e env.Environment) (driver.Conn, error) {
	dsn := fmt.Sprintf("%s:%s@%s:%d/%s", e.User, e.Password, e.Host, e.Port, e.Database)
	db, err := sql.Open("firebirdsql", dsn)
	if err != nil {
		return nil, err
	}

	return Conn{
		db:     db,
		schema: e.Database,
		table:  e.Table,
	}, nil
}

// Conn holds a connection to a FirebirdSQL database schema.
type Conn struct {
	db     *sql.DB
	schema string
	table  string
}

// HasTable check if the schema has the migration table needed for Rambler to operate on it.
func (c Conn) HasTable() (bool, error) {
	var table string
	err := c.db.QueryRow(`select table_name from information_schema.tables where table_schema = ? and table_name = ?`, c.schema, c.table).Scan(&table)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	if err == sql.ErrNoRows {
		return false, nil
	}

	return true, nil
}

// CreateTable create the migration table using a MySQL-compatible syntax.
func (c Conn) CreateTable() error {
	_, err := c.db.Exec(fmt.Sprintf(`CREATE TABLE %s ( migration VARCHAR(255) NOT NULL, PRIMARY KEY(migration) ) DEFAULT CHARSET=utf8`, c.table))
	return err
}

// GetApplied returns the list of already applied migrations.
func (c Conn) GetApplied() ([]string, error) {
	rows, err := c.db.Query(fmt.Sprintf(`SELECT migration FROM %s ORDER BY migration ASC`, c.table))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var migrations []string
	for rows.Next() {
		var migration string
		err := rows.Scan(&migration)
		if err != nil {
			return nil, err
		}

		migrations = append(migrations, migration)
	}

	return migrations, nil
}

// AddApplied record that a migration was applied.
func (c Conn) AddApplied(migration string) error {
	_, err := c.db.Exec(fmt.Sprintf(`INSERT INTO %s (migration) VALUES (?)`, c.table), migration)
	return err
}

// RemoveApplied record that a migration was reversed.
func (c Conn) RemoveApplied(migration string) error {
	_, err := c.db.Exec(fmt.Sprintf(`DELETE FROM %s WHERE migration = ?`, c.table), migration)
	return err
}

// Execute run a statement on the schema.
func (c Conn) Execute(statement string) error {
	_, err := c.db.Exec(statement)
	return err
}
