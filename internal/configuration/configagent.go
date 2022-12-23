package configuration

import (
	"log"
	"time"

	"github.com/caarlos0/env/v6"
)

type AgentConfig struct {
	Address        string        `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
	PollInterval   time.Duration `env:"POLL_INTERVAL" envDefault:"2s"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL" envDefault:"10s"`
}

func (c *AgentConfig) ToString() string {
	config := "Agent Config Settings: \n"
	config += "+ Address: " + c.Address + "\n"
	config += "+ PollInterval: " + c.PollInterval.String() + "\n"
	config += "+ ReportInterval: " + c.ReportInterval.String() + "\n"
	return config
}

func GetAgentConfig() *AgentConfig {
	cfg := &AgentConfig{}

	if err := env.Parse(cfg); err != nil {
		log.Println("Error during parsing the settings from env:")
		log.Printf(err.Error() + "\n\n")
	}

	log.Println(cfg.ToString())
	return cfg
}
