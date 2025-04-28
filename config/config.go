package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

type Config struct {
	DBhost        string
	DBport        string
	DBuser        string
	DBpassword    string
	DBname        string
	Port          string
	JWTSecret     string
	JWTExpiration int64
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot load env file")
	}

	config := &Config{
		DBhost:     getEnv("DB_PORT"),
		DBport:     getEnv("DB_PORT"),
		DBuser:     getEnv("DB_USER"),
		DBpassword: getEnv("DB_PASSWORD"),
		DBname:     getEnv("DB_NAME"),
		Port:       getEnv("PORT"),
		JWTSecret:  getEnv("JWT_SECRET"),
	}
	jwtExpiration, err := strconv.ParseInt(os.Getenv("JWT_EXPIRATION"), 10, 64)
	if err != nil {
		jwtExpiration = 3600
	}

	config.JWTExpiration = jwtExpiration
	return config
}

/*
@Paramater: c *Config -> receiver cho phep truy cap struct Config
@Return: *gorm.DB -> doi tuong de gorm thao tac voi csdl
  - error
*/
func (c *Config) ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", // data source name
		c.DBhost, c.DBport, c.DBuser, c.DBpassword, c.DBname)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func getEnv(key string) string {
	return os.Getenv(key)
}
