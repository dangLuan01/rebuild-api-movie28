package config

import (
	"fmt"
	"os"

	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
)

type DatabaseConfig struct {
	Host string
	Port string
	User string
	Password string
	DBName string
	SSLMode string
}
type RedisConfig struct {
    Host     string
    Port     string
    Password string
    DB       string
}
type Config struct {
	ServerAddress string
	DB DatabaseConfig
	Redis RedisConfig
}

func NewConfig() *Config {
	return &Config{
		ServerAddress: fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")),
		DB: DatabaseConfig{
			Host: utils.GetEnv("DB_HOST","localhost"),
			Port: utils.GetEnv("DB_PORT","3306"),
			User: utils.GetEnv("DB_USER","root"),
			Password: utils.GetEnv("DB_PASSWORD",""),
			DBName: utils.GetEnv("DB_DBNAME","mysql"),
			SSLMode: utils.GetEnv("DB_SSLMODE","disable"),
		},
		Redis: RedisConfig{
			Host: utils.GetEnv("REDIS_HOST", "localhost"),
			Port: utils.GetEnv("REDIS_PORT", "6379"),
			Password: utils.GetEnv("REDIS_PASSWORD", ""),
			DB: utils.GetEnv("REDIS_DB", "0"),
		},
	}
}

func (c *Config) DNS() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
    	c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.DBName,
	)
}