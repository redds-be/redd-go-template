package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

func (s *DBSuite) TestDB() {
	// Connect to the database
	frameRight, err := NewAdapter("sqlite3", "test_database.db")
	s.Require().NoError(err)

	// Create a table
	err = frameRight.CreateTable()
	s.Require().NoError(err)

	// Add an entry to the db
	err = frameRight.AddToHistory("Hello, World!")
	s.Require().NoError(err)

	// Close database connection
	err = frameRight.CloseDBConnection()
	s.Require().NoError(err)

	// Remove the database
	err = os.Remove("test_database.db")
	s.Require().NoError(err)
}

type DBSuite struct {
	suite.Suite
}

func TestDBSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(DBSuite))
}
