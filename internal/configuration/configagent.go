package configuration

import (
	"flag"
	"log"
	"os"
	"time"
)

type AgentConfig struct {
	Address        string        `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
	PollInterval   time.Duration `env:"POLL_INTERVAL" envDefault:"2s"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL" envDefault:"10s"`
	HashKey        string        `env:"KEY" envDefault:""`
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

	address := flag.String("a", "127.0.0.1:8080", "address for server listen")
	pollInterval := flag.Duration("p", time.Second*2, "interval for pollilng frequency")
	reportInterval := flag.Duration("r", time.Second*10, "interval for reporting frequency")
	hashKey := flag.String("k", "", "hash key for signing the data")
	flag.Parse()

	cfg.Address = *address
	cfg.PollInterval = *pollInterval
	cfg.ReportInterval = *reportInterval
	cfg.HashKey = *hashKey

	if addressFromEnv, ok := os.LookupEnv("ADDRESS"); ok {
		cfg.Address = addressFromEnv
	}

	if pollIntervalFromEnv, ok := os.LookupEnv("POLL_INTERVAL"); ok {
		dur, err := time.ParseDuration(pollIntervalFromEnv)
		if err != nil {
			log.Println("Couldn't parse the Duration from POLL_INTERVAL correctly:")
			log.Printf(err.Error() + "\n\n")
		}
		cfg.PollInterval = dur
	}

	if reportIntervalFromEnv, ok := os.LookupEnv("REPORT_INTERVAL"); ok {
		dur, err := time.ParseDuration(reportIntervalFromEnv)
		if err != nil {
			log.Println("Couldn't parse the Duration from REPORT_INTERVAL correctly:")
			log.Printf(err.Error() + "\n\n")
		}
		cfg.ReportInterval = dur
	}

	if hashKeyFromEnv, ok := os.LookupEnv("KEY"); ok {
		cfg.HashKey = hashKeyFromEnv
	}

	/* Legacy for reading env variables
	if err := env.Parse(cfg); err != nil {
		log.Println("Error during parsing the settings from env:")
		log.Printf(err.Error() + "\n\n")
	}*/

	log.Println(cfg.ToString())
	return cfg
}
