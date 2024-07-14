package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server Server
	Logger Logger
	Tuya   Tuya
}

type Server struct {
	Port  string
	Mode  string
	Debug bool
}

type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

type Tuya struct {
	Host     string
	ClientId string
	Secret   string
}

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil

}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("Unable to decode into sturct, %v", err)
		return nil, err
	}

	return &c, nil
}
