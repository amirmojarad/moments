package conf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadDbFileEnv(t *testing.T) {
	err := LoadDbEnv()
	assert.True(t, err == nil)
}
