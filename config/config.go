package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port          string `env:"PORT"`
	StanClusterId string `env:"STAN_CLUSTER_ID"`
	ClientId      string `env:"CLIENT_ID"`
	Subject       string `env:"SUBJECT"`
	DurableName   string `env:"DURABLE_NAME"`
	DSN           string `env:"DSN"`
	DriverName    string `env:"DRIVER_NAME"`
	NatsUrl       string `env:"NATS_URL"`
}

func GetConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	return &Config{
		Port:          os.Getenv("PORT"),
		StanClusterId: os.Getenv("STAN_CLUSTER_ID"),
		ClientId:      os.Getenv("CLIENT_ID"),
		Subject:       os.Getenv("SUBJECT"),
		DurableName:   os.Getenv("DURABLE_NAME"),
		DSN:           os.Getenv("DSN"),
		DriverName:    os.Getenv("DRIVER_NAME"),
		NatsUrl:       os.Getenv("NATS_URL"),
	}, nil

}
