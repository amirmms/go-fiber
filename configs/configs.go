package configs

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"log"
)

const (
	PROD = "production"
	DEV  = "develop"
)

type Config struct {
	AppEnv  string `env:"APP_ENV" envDefault:"local"`
	AppName string `env:"APP_NAME"`
	AppHost string `env:"APP_HOST" envDefault:""`
	AppPort string `env:"APP_PORT" envDefault:"8080"`

	DBHost     string `env:"DB_HOST" envDefault:"localhost"`
	DBPort     string `env:"DB_PORT" envDefault:"5432"`
	DBDatabase string `env:"DB_DATABASE" envDefault:"database"`
	DBUsername string `env:"DB_USERNAME" envDefault:"root"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"secret"`

	RedisHost     string `env:"REDIS_HOST" envDefault:"localhost"`
	RedisPort     string `env:"REDIS_PORT" envDefault:"6379"`
	RedisPassword string `env:"REDIS_PASSWORD"`

	BodyLimit          int `env:"BODY_LIMIT" envDefault:"4198400"`
	MaxConnsPerIP      int `env:"MAX_CONNS_PER_IP"`
	MaxRequestsPerConn int `env:"MAX_REQUESTS_PER_CONN"`

	EncryptCookieKey string `env:"ENCRYPT_COOKIE_KEY"`
}

func Init() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	conf := Config{}

	err = env.Parse(&conf)
	if err != nil {
		log.Fatal(err)
	}

	return conf
}
