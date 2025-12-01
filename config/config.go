package config

import (
	"log"
	"sync"

	"github.com/furqanrupom/sqlc-crud/util"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host          string
	DBPort        int
	User          string
	Password      string
	DBName        string
	SslEnableMode string
}

type Config struct {
	HttpPort int
	DB       DBConfig
}

var (
	cfg  *Config
	once sync.Once
)

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env load failed!")
	}

	cfg = &Config{
		HttpPort: util.CheckEnvInt("PORT"),
		DB: DBConfig{
			Host:          util.CheckEnvStr("HOST"),
			User:          util.CheckEnvStr("DB_USER"),
			Password:      util.CheckEnvStr("PASSWORD"),
			DBPort:        util.CheckEnvInt("DB_PORT"),
			DBName:        util.CheckEnvStr("DB_NAME"),
			SslEnableMode: util.CheckEnvStr("ENABLE_SSL_MODE"),
		},
	}
}

func GetConfig() *Config {
	once.Do(loadConfig)
	return cfg
}
