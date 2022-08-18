package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type EnvVariables struct {
	Topic         string
	BrokerAddress string
	Partition     int
	RouterPort    string
	GroupId       string
}

var env *EnvVariables

func initEnvVariables() {
	if env != nil {
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading during .env file")
	}

	partition, err := strconv.Atoi(os.Getenv("PARTITION"))

	if err != nil {
		log.Fatal("Can not convert partition value string to int")
	}

	env = &EnvVariables{
		Topic:         os.Getenv("TOPIC"),
		BrokerAddress: os.Getenv("BROKER_ADDRESS"),
		Partition:     partition,
		RouterPort:    os.Getenv("ROUTER_PORT"),
		GroupId:       os.Getenv("GROUP_ID"),
	}
}
func GetEnv() *EnvVariables {
	initEnvVariables()
	if env == nil {
		log.Fatal("Error occurred during initEnvVariable method.")
	}
	return env
}
