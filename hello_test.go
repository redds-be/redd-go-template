package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func (s *HelloSuite) TestHelloWorld() {
	// Test the helloWorld() function
	expectedHelloWorld := "Hello, World!"
	actualHelloWorld := helloWorld()

	// Is the message returned by helloWorld() corresponding to the expected message?
	s.Equal(expectedHelloWorld, actualHelloWorld)
}

type HelloSuite struct {
	suite.Suite
}

func TestHelloSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(HelloSuite))
}
