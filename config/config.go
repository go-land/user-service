package config

import "os"

type EnvConfig struct {
	Port    string
	Profile string
}

func NewEnvConfig() *EnvConfig {
	return &EnvConfig{
		Port:    os.Getenv("PORT"),
		Profile: os.Getenv("PROFILE"),
	}
}
