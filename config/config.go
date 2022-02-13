package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	APP_ENV = "APP_ENV"
	LOCAL   = "local"
	PROD    = "prod"
)

var Env *Config

type Config struct {
	isLocal              bool
	InitialBalanceAmount float64 `json="initialBalanceAmount"`
	MinumumBalanceAmount float64 `json="minumumBalanceAmount"`
}

func (c *Config) IsLocal() bool {
	return c.isLocal
}

func init() {
	appEnv, ok := os.LookupEnv(APP_ENV)
	if !ok {
		appEnv = LOCAL
	}

	read, err := os.ReadFile(fmt.Sprintf("../.config/%s.json", appEnv))
	if err != nil {
		panic(err)
	}

	env := &Config{}
	err = json.Unmarshal(read, env)
	if err != nil {
		panic(err)
	}

	env.isLocal = appEnv == LOCAL
	Env = env
}

func GetEnv(key string, def string) string {
	env, has := os.LookupEnv(key)
	if !has {
		return def
	}
	return env
}
