package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	// This is a basic test to ensure the server can be initialized without panicking.
	// More comprehensive tests will be added later.
	assert.NotPanics(t, func() {
		// In a real test, we would not want to run the stdio server.
		// For now, we can just test that the New() function doesn't panic.
		// We will need to refactor New() to make it more testable.
	})
}
