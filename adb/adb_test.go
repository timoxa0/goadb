package adb

import (
	"testing"

	"github.com/timoxa0/goadb/wire"

	"github.com/stretchr/testify/assert"
)

func TestGetServerVersion(t *testing.T) {
	s := &MockServer{
		Status:   wire.StatusSuccess,
		Messages: []string{"000a"},
	}
	client := &Adb{s}

	v, err := client.ServerVersion()
	assert.Equal(t, "host:version", s.Requests[0])
	assert.NoError(t, err)
	assert.Equal(t, 10, v)
}
