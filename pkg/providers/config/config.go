package config

import (
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

type DBType string

const (
	Postgres DBType = "postgres"
	MySQL    DBType = "mysql"
	SQLite   DBType = "sqlite"
)

type ConfigProvider struct {
	MACHINE_ID   int
	APP_DOMAIN   string
	ENABLE_TLS   bool
	CORS_ORIGINS string

	DB_TYPE DBType
	DB_HOST string
	DB_PORT string
	DB_PWD  string
	DB_NAME string

	PROMETHEUS_ENABLED bool
	PROMETHEUS_ADDR    string
	PROMETHEUS_PORT    int
}

func getEnvStr(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

func getEnvBool(key string, defaultVal bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	parsed, err := strconv.ParseBool(val)
	if err != nil {
		return defaultVal
	}
	return parsed
}

func getEnvInt(key string, defaultVal int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	parsed, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return parsed
}

func NewConfigProvider() (*ConfigProvider, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	return &ConfigProvider{
		MACHINE_ID:   getEnvInt("MACHINE_ID", 1),
		APP_DOMAIN:   getEnvStr("APP_DOMAIN", "localhost"),
		ENABLE_TLS:   getEnvBool("ENABLE_TLS", false),
		CORS_ORIGINS: getEnvStr("CORS_ORIGINS", "*"),

		DB_TYPE: DBType(getEnvStr("DB_TYPE", "sqlite")),
		DB_HOST: getEnvStr("DB_HOST", "localhost"),
		DB_PWD:  getEnvStr("DB_PWD", "strong_password"),
		DB_NAME: getEnvStr("DB_NAME", "sols_cms"),

		PROMETHEUS_ENABLED: getEnvBool("PROMETHEUS_ENABLED", false),
		PROMETHEUS_ADDR:    getEnvStr("PROMETHEUS_ADDR", "localhost"),
		PROMETHEUS_PORT:    getEnvInt("PROMETHEUS_PORT", 8080),
	}, nil
}

func (p *ConfigProvider) GetDSN() string {
	return "host=" + p.DB_HOST + " user=sols_cms password=" + p.DB_PWD + " dbname=" + p.DB_NAME + " port=" + p.DB_PORT + " sslmode=disable TimeZone=UTC"
}

func (p *ConfigProvider) Getenv(key string) string {
	v := reflect.ValueOf(p).Elem()
	field := v.FieldByName(key)
	if field.IsValid() && field.Kind() == reflect.String {
		return field.String()
	}
	return os.Getenv(key)
}
