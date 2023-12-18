package ports

// DBPort is the port for a database adapter.
type DBPort interface {
	CloseDBConnection() error
	CreateTable() error
	AddToHistory(message string) error
}
