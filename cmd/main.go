package main

import (
	"log"

	"github.com/redds-be/redd-go-template/internal/application/api"
	"github.com/redds-be/redd-go-template/internal/application/core/helloworld"
	"github.com/redds-be/redd-go-template/internal/framework/in/http"
	"github.com/redds-be/redd-go-template/internal/framework/out/database"
)

// main drives the application.
func main() {
	var err error

	// Create the database
	dbAdapter, err := database.NewAdapter("sqlite3", "hello.db")
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %s", err)
	}

	// Close the database at the end
	defer func(dbAdapter *database.Adapter) {
		err := dbAdapter.CloseDBConnection()
		if err != nil {
			log.Fatalf("failed to close the database: %s", err)
		}
	}(dbAdapter)

	// Create the default table
	err = dbAdapter.CreateTable()
	if err != nil {
		log.Panicf("failed to create the table: %s", err)
	}

	// helloworld
	hello := helloworld.New()
	log.Println(hello.HelloWorld())

	// Initiate the api
	applicationAPI := api.NewApplication(dbAdapter, hello)

	// Initiate the adapter using the api
	httpAdapter := http.NewAdapter(applicationAPI)

	// Start the http server that will serve a hello world message on http://localhost:8080/
	log.Println("Listening on port 8080 at localhost")
	err = httpAdapter.Run()
	if err != nil {
		log.Panicf("could not listen on port 8080 at localhost: %s", err)
	}
}
