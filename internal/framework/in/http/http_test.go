package http

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/redds-be/redd-go-template/internal/application/api"
	"github.com/redds-be/redd-go-template/internal/application/core/helloworld"
	"github.com/redds-be/redd-go-template/internal/framework/out/database"
	"github.com/stretchr/testify/suite"
)

func (s *HTTPSuite) TestAddition() {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://localhost:8080/", nil)
	s.Require().NoError(err)
	resp := httptest.NewRecorder()
	s.httpa.hello(resp, req)
	s.Equal(http.StatusOK, resp.Code)
	s.Equal("<h1>Hello, World!</h1>", resp.Body.String())
}

type HTTPSuite struct {
	httpa Adapter
	suite.Suite
}

func TestHTTPSuite(t *testing.T) {
	dbAdapter, err := database.NewAdapter("sqlite3", "test_http.db")
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
			t.Fatal("unable to close the database connection")
		}
	}(dbAdapter)

	defer func() {
		err := os.Remove("test_http.db")
		if err != nil {
			t.Fatal("unable to remove the test database")
		}
	}()

	hello := helloworld.New()
	applicationAPI := api.NewApplication(dbAdapter, hello)

	t.Parallel()
	suite.Run(t, &HTTPSuite{
		httpa: Adapter{api: applicationAPI},
		Suite: suite.Suite{},
	})
}
