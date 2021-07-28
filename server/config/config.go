package config

import (
	"github.com/joho/godotenv"
	"github.com/tkanos/gonfig"
	"log"
)

type ServerConfig struct {
	Hash HashConfig
	Host HostConfig
	Database DBConfig
}

type HashConfig struct {
	NumOfHashSymbols int `env:"NUM_OF_HASH_SYMBOLS"`
	HashSymbols string `env:"HASH_SYMBOLS"`
}

type HostConfig struct {
	ServerHost string `env:"SERVER_HOST"`
	ServerStartPort string `env:"SERVER_START_PORT"`
}

type DBConfig struct {
	DBHost string `env:"DB_HOST"`
	DBPort int `env:"DB_PORT"`
	DBUser string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName string `env:"DB_NAME"`
}

func NewServerConfig() *ServerConfig {
	var err error

	var hashConfig HashConfig
	err = getConfig(&hashConfig)
	if err != nil {
		log.Println(err)
	}

	var hostConfig HostConfig
	err = getConfig(&hostConfig)
	if err != nil {
		log.Println(err)
	}

	var dbConfig DBConfig
	err = getConfig(&dbConfig)
	if err != nil {
		log.Println(err)
	}

	return &ServerConfig{
		Hash: hashConfig,
		Host: hostConfig,
		Database: dbConfig,
	}
}

func getConfig(configType interface{}) (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
	return gonfig.GetConf("", configType)
}