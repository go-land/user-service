package config

type EnvConfig struct {
	Port    string
	Profile string
}

func NewEnvConfig() *EnvConfig {
	return &EnvConfig{
		Port:    "7070", //os.Getenv("PORT"),
		Profile: "dev",  //os.Getenv("PROFILE"),
	}
}
