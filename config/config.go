package config

import (
	"os"
	"sync"

	"github.com/BurntSushi/toml"
)

var (
	configuration *Config
	configLock    = new(sync.RWMutex)
)

type Config struct {
	HTTP        HTTP        `toml:"http"`
	ECS         ECS         `toml:"ecs"`
	Credentials Credentials `toml:"credentials"`
}

type HTTP struct {
	Port string `toml:"port"`
}

type ECS struct {
	ClusterName string `toml:"cluster_name"`
}

type Credentials struct {
	AccessKeyID     string `toml:"access_key_id"`
	SecretAccessKey string `toml:"secret_access_key"`
}

func Get() *Config {
	configLock.RLock()
	defer configLock.RUnlock()
	return configuration
}

func Load(path string) error {
	configLock.Lock()
	defer configLock.Unlock()

	configuration = &Config{
		HTTP: HTTP{
			Port: os.Getenv("PORT"),
		},

		ECS: ECS{
			ClusterName: os.Getenv("CLUSTER_NAME"),
		},

		Credentials: Credentials{
			AccessKeyID:     os.Getenv("ACCESS_KEY_ID"),
			SecretAccessKey: os.Getenv("SECRET_ACCESS_KEY"),
		},
	}

	_, err := toml.DecodeFile(path, &configuration)
	return err
}
