package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/redds-be/redd-go-template/internal/ports"
)

// Adapter is a rpc adapter the is compatible with the api.
type Adapter struct {
	api ports.APIPort
}

// NewAdapter creates a grpc entry adapter that is compatible with gRPCEntryPort.
func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

// Run is the HTTP entrypoint for the application.
func (httpa Adapter) Run() error {
	http.HandleFunc("/", httpa.hello)

	// Listen and serve
	err := http.ListenAndServe(fmt.Sprintf(":%s", "8080"), nil) //nolint:gosec

	return err
}

// hello tells hello world to a http client.
func (httpa Adapter) hello(writer http.ResponseWriter, _ *http.Request) {
	message, err := httpa.api.GetHelloWorld()
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	writer.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(writer, "<h1>%s</h1>", message)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
}
