package api

import (
	"log"
	"os"
	"testing"

	"github.com/redds-be/redd-go-template/internal/application/core/helloworld"
	"github.com/redds-be/redd-go-template/internal/framework/out/database"
	"github.com/stretchr/testify/suite"
)

func (s *APISuite) TestGetHelloWorld() {
	message, err := s.testApp.GetHelloWorld()
	s.Require().NoError(err)
	s.Equal("Hello, World!", message)
}

type APISuite struct {
	testApp Application
	suite.Suite
}

func TestAPISuite(t *testing.T) {
	dbAdapter, err := database.NewAdapter("sqlite3", "test_api.db")
	if err != nil {
		log.Fatal(err)
	}

	err = dbAdapter.CreateTable()
	if err != nil {
		log.Fatal(err)
	}

	defer func(dbAdapter *database.Adapter) {
		err := dbAdapter.CloseDBConnection()
		if err != nil {
			t.Fatal("unable to close the database")
		}
	}(dbAdapter)

	defer func() {
		err := os.Remove("test_api.db")
		if err != nil {
			t.Fatal("unable to remove the test database")
		}
	}()

	core := helloworld.New()

	t.Parallel()
	suite.Run(t, &APISuite{
		testApp: Application{dbase: dbAdapter, hello: core},
		Suite:   suite.Suite{},
	})
}
