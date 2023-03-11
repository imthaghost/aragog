package config

// Service is an interface that defines the functions needed to implement a Config Service.
type Service interface {
	// Load will do any config setup (like load env vars)
	Load()
	// Get will get the config
	Get() Config
}

// Config is a service that is designed to provide various configuration to the rest of the application.
type Config struct {
	General  GeneralConfig
	Datadog  DatadogConfig
	RabbitMQ RabbitMQConfig
	Monopoly MonopolyClientConfig
	Sentry   SentryConfig
}

// GeneralConfig contains general information that the service needs to run.
type GeneralConfig struct {
	SiteURL       string // application URL
	AppEnv        string // the environment that the application is running in (dev, prod, etc)
	SessionSecret string // session secret for cookies (set in env var)
}

// DatadogConfig keeps track of datadog configuration
type DatadogConfig struct {
	AgentURL string
}

// RabbitMQConfig ...
type RabbitMQConfig struct {
	Host     string
	Username string
	Password string
}

// MonopolyClientConfig keeps track of Monopoly Client configuration
type MonopolyClientConfig struct {
	Host string
}

// SentryConfig contains information that allows us to interact with the Sentry API
type SentryConfig struct {
	DSN string
}

// getSentryConfig returns the Sentry config
func getSentryConfig() SentryConfig {
	env := getGeneralConfig().AppEnv

	config := SentryConfig{
		DSN: "https://11537d124de74bfbb7bf2febe4148287@o1283698.ingest.sentry.io/6591504",
	}

	if env == STAGING {
		config.DSN = "https://a20f62edff0d4aa98a0be41c3af33145@o1283698.ingest.sentry.io/6591506"
	}

	if env == PROD {
		config.DSN = "https://b268f04e5bf045dab71629b131c8dd35@o1283698.ingest.sentry.io/6591508"
	}

	return config
}
