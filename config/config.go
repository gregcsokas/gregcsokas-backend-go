package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	// Database config
	DBHost     string `env:"DB_HOST" envDefault:"localhost"`
	DBPort     string `env:"DB_PORT" envDefault:"5432"`
	DBUser     string `env:"DB_USER" envDefault:"postgres"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"postgres"`
	DBName     string `env:"DB_NAME" envDefault:"postgres"`
	DBSSLMode  string `env:"DB_SSL_MODE" envDefault:"disable"`
	DBTimeZone string `env:"DB_TIME_ZONE" envDefault:"Etc/Universal"`

	// Server configs
	ServerPort string `env:"SERVER_PORT" envDefault:"8060"`
	ServerEnv  string `env:"SERVER_ENV" envDefault:"development"`

	SecretKey    string `env:"SECRET_KEY" envDefault:"sup3rs3cr3t"`
	AllowedHosts string `env:"ALLOWED_HOSTS" envDefault:"*"`
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	config := &Config{
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		DBSSLMode:    os.Getenv("DB_SSL_MODE"),
		DBTimeZone:   os.Getenv("DB_TIME_ZONE"),
		ServerPort:   os.Getenv("SERVER_PORT"),
		ServerEnv:    os.Getenv("SERVER_ENV"),
		SecretKey:    os.Getenv("SECRET_KEY"),
		AllowedHosts: os.Getenv("ALLOWED_HOSTS"),
	}

	if err := config.Validate(); err != nil {
		log.Fatal("Fatal error when validating config.")
		return nil, err
	}

	return config, nil
}

func (c *Config) Validate() error {
	required := map[string]string{
		"DB_NAME":     c.DBName,
		"DB_USER":     c.DBUser,
		"DB_PASSWORD": c.DBPassword,
		"SECRET_KEY":  c.SecretKey,
	}

	for k, v := range required {
		if v == "" {
			return fmt.Errorf("Required field '%s' is missing", k)
		}
	}

	return nil
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s timezone=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode, c.DBTimeZone,
	)
}
