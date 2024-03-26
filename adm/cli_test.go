package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/plumk97/go-admin/modules/system"
)

func TestGetLatestVersion(t *testing.T) {
	assert.Equal(t, getLatestVersion(), system.Version())
}
