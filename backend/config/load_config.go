package config

import (
	"github.com/Netflix/go-env"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST" env:"DB_HOST"`
	DBPort     int    `mapstructure:"DB_PORT" env:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER" env:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD" env:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_TABLE" env:"DB_TABLE"`
	JWTSecret  string `mapstructure:"JWT_SECRET" env:"JWT_SECRET"`
}

func LoadConfigFromFile(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}

func LoadConfigFromEnv() (config Config, err error) {
	_, err = env.UnmarshalFromEnviron(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
