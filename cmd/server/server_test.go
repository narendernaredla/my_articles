package server

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartServerWithoutEnv(t *testing.T) {
	server := NewServer()
	expectedAddr := ":8080"
	assert.NotNil(t, server)
	assert.Equal(t, expectedAddr, server.Addr)
}

func TestStartServerWithEnv(t *testing.T) {
	os.Setenv("PORT", "8000")
	server := NewServer()
	expectedAddr := ":8000"
	assert.NotNil(t, server)
	assert.Equal(t, expectedAddr, server.Addr)
}
