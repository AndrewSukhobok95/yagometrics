package configuration

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type ServerConfig struct {
	Address       string        `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
	StoreInterval time.Duration `env:"STORE_INTERVAL" envDefault:"300s"`
	StoreFile     string        `env:"STORE_FILE" envDefault:"/tmp/devops-metrics-db.json"`
	Restore       bool          `env:"RESTORE" envDefault:"true"`
	HashKey       string        `env:"KEY" envDefault:""`
	DBAddress     string        `env:"DATABASE_DSN" envDefault:""`
}

func (c *ServerConfig) ToString() string {
	config := "Server Config Settings: \n"
	config += "+ Address: " + c.Address + "\n"
	config += "+ StoreFile: " + c.StoreFile + "\n"
	config += "+ StoreInterval: " + c.StoreInterval.String() + "\n"
	config += "+ Restore: " + fmt.Sprintf("%t", c.Restore) + "\n"
	config += "+ DB Address: " + c.DBAddress + "\n"
	return config
}

func GetServerConfig() *ServerConfig {
	cfg := &ServerConfig{}

	address := flag.String("a", "127.0.0.1:8080", "address for server listen")
	restore := flag.String("r", "true", "restore latest values")
	storeFile := flag.String("f", "/tmp/devops-metrics-db.json", "file for db")
	storeInterval := flag.Duration("i", time.Second*300, "interval for db update")
	hashKey := flag.String("k", "", "hash key for signing the data")
	dbAddress := flag.String("d", "", "db address")
	flag.Parse()

	cfg.Address = *address
	switch *restore {
	case "false":
		cfg.Restore = false
	case "true":
		cfg.Restore = true
	default:
		cfg.Restore = true
	}
	cfg.StoreFile = *storeFile
	cfg.StoreInterval = *storeInterval
	cfg.HashKey = *hashKey
	cfg.DBAddress = *dbAddress

	if addressFromEnv, ok := os.LookupEnv("ADDRESS"); ok {
		cfg.Address = addressFromEnv
	}

	if restoreFromEnv, ok := os.LookupEnv("RESTORE"); ok {
		switch restoreFromEnv {
		case "false":
			cfg.Restore = false
		case "true":
			cfg.Restore = true
		default:
			cfg.Restore = true
		}
	}

	if storeIntervalFromEnv, ok := os.LookupEnv("STORE_INTERVAL"); ok {
		dur, err := time.ParseDuration(storeIntervalFromEnv)
		if err != nil {
			log.Println("Couldn't parse the Duration from STORE_INTERVAL correctly:")
			log.Printf(err.Error() + "\n\n")
		}
		cfg.StoreInterval = dur
	}

	if storeFileFromEnv, ok := os.LookupEnv("STORE_FILE"); ok {
		cfg.StoreFile = storeFileFromEnv
	}

	if hashKeyFromEnv, ok := os.LookupEnv("KEY"); ok {
		cfg.HashKey = hashKeyFromEnv
	}

	if dbAddressFromEnv, ok := os.LookupEnv("DATABASE_DSN"); ok {
		cfg.DBAddress = dbAddressFromEnv
	}

	/* Legacy for reading env variables
	if err := env.Parse(cfg); err != nil {
		log.Println("Error during parsing the settings from env:")
		log.Printf(err.Error() + "\n\n")
	}*/

	log.Println(cfg.ToString())
	return cfg
}
