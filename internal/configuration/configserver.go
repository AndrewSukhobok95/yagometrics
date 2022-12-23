package configuration

import (
	"fmt"
	"log"
	"time"

	"github.com/caarlos0/env/v6"
)

type ServerConfig struct {
	Address       string        `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
	StoreInterval time.Duration `env:"STORE_INTERVAL" envDefault:"300s"`
	StoreFile     string        `env:"STORE_FILE" envDefault:"/tmp/devops-metrics-db.json"`
	Restore       bool          `env:"RESTORE" envDefault:"true"`
}

func (c *ServerConfig) ToString() string {
	config := "Server Config Settings: \n"
	config += "+ Address: " + c.Address + "\n"
	config += "+ StoreFile: " + c.StoreFile + "\n"
	config += "+ StoreInterval: " + c.StoreInterval.String() + "\n"
	config += "+ Restore: " + fmt.Sprintf("%t", c.Restore) + "\n"
	return config
}

func GetServerConfig() *ServerConfig {
	cfg := &ServerConfig{}

	if err := env.Parse(cfg); err != nil {
		log.Println("Error during parsing the settings from env:")
		log.Printf(err.Error() + "\n\n")
	}

	log.Println(cfg.ToString())
	return cfg
}
