package ports

// APIPort is the port for driving adapters.
type APIPort interface {
	GetHelloWorld() (string, error)
}
