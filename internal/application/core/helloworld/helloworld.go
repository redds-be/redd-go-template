package helloworld

// Hello implements the Helloworld interface.
type Hello struct{}

// New creates a new Hello.
func New() *Hello {
	return &Hello{}
}

// HelloWorld returns a hello world message.
func (hello Hello) HelloWorld() string {
	return "Hello, World!"
}
