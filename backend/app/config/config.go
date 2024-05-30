package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DevMode   bool
	DbUrl     string
	JwtSecret string
}

func NewConfig() (*Config, error) {

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read from config: %v\n", err)
		return nil, err
	}

	return &Config{
		DevMode:   viper.GetBool("DEV_MODE"),
		DbUrl:     viper.GetString("DATABASE_URL"),
		JwtSecret: viper.GetString("JWT_SECRET"),
	}, nil

}
