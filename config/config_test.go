package config_test

import (
	"bootcamp/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	KEY = "KEY"
	DEF = "DEF"
	VAL = "VAL"
)

func TestGetEnvReturnDefault(t *testing.T) {
	env := config.GetEnv(KEY, DEF)
	assert.Equal(t, env, DEF)
}

func TestGetEnvReturnEvrioment(t *testing.T) {
	os.Setenv(KEY, VAL)
	env := config.GetEnv(KEY, DEF)
	assert.Equal(t, env, VAL)
	os.Setenv(KEY, "")
}

func TestConfigIsLocalIfLocal(t *testing.T) {
	assert.True(t, config.Env.IsLocal())
}
