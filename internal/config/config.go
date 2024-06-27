package config

import (
	"github.com/spf13/viper"
	"log"
)

var GlobalConfig *Config

type Config struct {
	MySQL        MySQLConfig `mapstructure:"mysql"`
	ServerPort   int         `mapstructure:"server_port"`
	JWTSecretKey string      `mapstructure:"JWT_SECRET_KEY"`
}

type MySQLConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	GlobalConfig = &cfg
}
