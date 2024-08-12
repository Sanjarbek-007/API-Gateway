package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"

	_ "github.com/lib/pq"
)

type Config struct {
	HTTP_PORT         string
	GRPC_USER_PORT    string
	GRPC_PRODUCT_PORT string
	DB_HOST           string
	DB_PORT           string
	DB_USER           string
	DB_PASSWORD       string
	DB_NAME           string
	ACCESS_TOKEN      string
}

func Load() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}

	config := Config{}

	config.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", ":8080"))
	config.GRPC_USER_PORT = cast.ToString(coalesce("GRPC_USER_PORT", 50050))
	config.GRPC_PRODUCT_PORT = cast.ToString(coalesce("GRPC_PRODUCT_PORT", 50051))
	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToString(coalesce("DB_PORT", "5432"))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "1111"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "postgres"))
	config.ACCESS_TOKEN = cast.ToString(coalesce("ACCESS_TOKEN", "key_is_really_easy"))

	return config
}

func coalesce(env string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(env)
	if !exists {
		return defaultValue
	}
	return value
}

func CasbinEnforcer(logger *slog.Logger) (*casbin.Enforcer, error) {
	config := Load()
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
	config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_NAME, config.DB_PASSWORD)
	fmt.Println(conn)
	adapter, err := xormadapter.NewAdapter("postgres", conn)
	if err != nil {
		log.Println("error creating Casbin adapter", "error", err.Error())
		logger.Error("Error creating Casbin adapter", "error", err.Error())
		return nil, err
	}

	enforcer, err := casbin.NewEnforcer("config/model.conf", adapter)
	if err != nil {
		logger.Error("Error creating Casbin enforcer", "error", err.Error())
		log.Println("error creating Casbin enforcer", "error", err.Error())
		return nil, err
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		log.Println("error loading Casbin policy", "error", err.Error())
		logger.Error("Error loading Casbin policy", "error", err.Error())
		return nil, err
	}

	policies := [][]string{
		{"user", "/api/users", "GET"},
		{"admin", "/api/users", "POST"},
	}

	_, err = enforcer.AddPolicies(policies)
	if err != nil {
		log.Println("error adding Casbin policy", "error", err.Error())
		logger.Error("Error adding Casbin policy", "error", err.Error())
		return nil, err
	}

	err = enforcer.SavePolicy()
	if err != nil {
		log.Println("Error saving Casbin policy", "error", err.Error())
		logger.Error("Error saving Casbin policy", "error", err.Error())
		return nil, err
	}
	return enforcer, nil
}
