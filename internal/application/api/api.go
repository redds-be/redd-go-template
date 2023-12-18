package api

import "github.com/redds-be/redd-go-template/internal/ports"

type Application struct {
	dbase ports.DBPort
	hello Helloworld
}

// NewApplication creates a new Application.
func NewApplication(dbase ports.DBPort, hello Helloworld) *Application {
	return &Application{dbase: dbase, hello: hello}
}

// GetHelloWorld gets the hello world message.
func (apia Application) GetHelloWorld() (string, error) {
	message := apia.hello.HelloWorld()

	err := apia.dbase.AddToHistory(message)
	if err != nil {
		return "", err
	}

	return message, nil
}
