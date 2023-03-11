package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	DEV     = "dev"
	PROD    = "prod"
	STAGING = "staging"
)

type New struct{}

// Load will load the config
func (n *New) Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Println("did not load env vars from .env")
	}
}

// Get will return the default config
func (n *New) Get() Config {
	// get the config
	return Config{
		General:  getGeneralConfig(),
		Datadog:  getDatadogConfig(),
		RabbitMQ: getRabbitMQConfig(),
		Monopoly: getMonopolyClientConfig(),
		Sentry:   getSentryConfig(),
	}
}

// getGeneralConfig will return general configuration values
func getGeneralConfig() GeneralConfig {
	// default
	config := GeneralConfig{
		AppEnv:        os.Getenv("APP_ENV"),
		SiteURL:       "http://localhost:8080/",
		SessionSecret: os.Getenv("SESSION_SECRET"),
	}
	// beta
	if config.AppEnv == STAGING {
		config.SiteURL = "https://staging.aragog.scaletrade.ai/"
	}
	// prod
	if config.AppEnv == PROD {
		config.SiteURL = "https://aragog.scaletrade.ai/"
	}

	return config
}

// getDatadogConfig will return the default Datadog config
func getDatadogConfig() DatadogConfig {
	return DatadogConfig{
		AgentURL: os.Getenv("DATADOG_AGENT_URL"),
	}
}

// getRabbitMQConfig ...
func getRabbitMQConfig() RabbitMQConfig {

	return RabbitMQConfig{
		Host:     os.Getenv("RABBITMQ_HOST"),
		Username: os.Getenv("RABBITMQ_USERNAME"),
		Password: os.Getenv("RABBITMQ_PASSWORD"),
	}

}

func getMonopolyClientConfig() MonopolyClientConfig {
	env := getGeneralConfig().AppEnv

	config := MonopolyClientConfig{
		Host: "http://localhost:8080/",
	}

	if env == STAGING {
		config.Host = "https://staging.scaletrade.ai/"
	}

	if env == PROD {
		config.Host = "https://scaletrade.ai/"
	}
	return config
}
