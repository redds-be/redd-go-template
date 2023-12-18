package helloworld

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func (s *HelloSuite) TestHello() {
	adapter := New()
	message := adapter.HelloWorld()
	s.Equal("Hello, World!", message)
}

type HelloSuite struct {
	suite.Suite
}

func TestHelloSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(HelloSuite))
}
