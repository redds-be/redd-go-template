package database

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3" // Driver for sqlite
)

// Adapter implements the DBPort interface.
type Adapter struct {
	dbase *sql.DB
}

// NewAdapter creates a new Adapter.
func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	// Connect to the database
	dbase, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	// Test db connection
	err = dbase.Ping()
	if err != nil {
		return nil, err
	}

	return &Adapter{dbase: dbase}, nil
}

// CloseDBConnection closes the db connection.
func (dbAdapter Adapter) CloseDBConnection() error {
	// Close the database
	err := dbAdapter.dbase.Close()

	return err
}

// CreateTable create the default 'hello_history' table.
func (dbAdapter Adapter) CreateTable() error {
	const queryString = `CREATE TABLE IF NOT EXISTS hello_history (date TIMESTAMP, message VARCHAR(64));`
	_, err := dbAdapter.dbase.Exec(queryString)

	return err
}

// AddToHistory adds the helloworld message to the database history table.
func (dbAdapter Adapter) AddToHistory(message string) error {
	const queryString = `INSERT INTO hello_history VALUES ($1, $2)`
	_, err := dbAdapter.dbase.Exec(queryString, time.Now().UTC(), message)

	return err
}
