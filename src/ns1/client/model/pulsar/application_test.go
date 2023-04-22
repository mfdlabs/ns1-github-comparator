package pulsar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewApplication(t *testing.T) {
	a := NewApplication("app_name")
	assert.Equal(t, "app_name", a.Name, "Wrong name")
}
