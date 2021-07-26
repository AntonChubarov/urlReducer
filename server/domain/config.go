package domain

import (
	"os"
	"strconv"
)

type ServerConfig struct {
	NumOfHashSymbols int
	HashRegex string
	URLRegex string
	ServerHost string
	ServerStartPort string
	DBHost string
	DBPort int
	DBUser string
	DBPassword string
	DBName string
}

func NewServerConfig() ServerConfig {
	return ServerConfig{
		NumOfHashSymbols: getEnvAsInt("NUM_OF_HASH_SYMBOLS", 7),
		HashRegex:        getEnvAsString("HASH_REGEX", "^[a-zA-Z0-9]*$"),
		URLRegex:         getEnvAsString("URL_REGEX", "https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)"),
		ServerHost:       getEnvAsString("SERVER_HOST", "http://localhost"),
		ServerStartPort:  getEnvAsString("SERVER_START_PORT", ":8080"),
		DBHost:           getEnvAsString("DB_HOST", "localhost"),
		DBPort:           getEnvAsInt("DB_PORT", 5432),
		DBUser:           getEnvAsString("DB_USER", "postgres"),
		DBPassword:       getEnvAsString("DB_PASSWORD", "Cc030789"),
		DBName:           getEnvAsString("DB_NAME", "postgres"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvAsString(name string, defaultVal string) string {
	valueStr := getEnv(name, defaultVal)
	return valueStr
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}