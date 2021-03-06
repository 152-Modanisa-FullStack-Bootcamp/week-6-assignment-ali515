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

func Initialize() {
	appEnv, ok := os.LookupEnv(APP_ENV)
	if !ok {
		appEnv = LOCAL
	}

	path := fmt.Sprintf(".config/%s.json", appEnv)
	read, err := os.ReadFile(path)
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
