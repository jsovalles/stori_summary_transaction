package config

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Config struct {
	Environment  string `mapstructure:"environment"`
	SmtpServer   string `mapstructure:"smtp-server"`
	SmtpPort     string `mapstructure:"smtp-port"`
	SmtpUsername string `mapstructure:"smtp-username"`
	SmtpPassword string `mapstructure:"smtp-password"`
	GeneralEmail string `mapstructure:"general-email"`
	DbHost       string `mapstructure:"db-host"`
	DbPort       string `mapstructure:"db-port"`
	DbUsername   string `mapstructure:"db-username"`
	DbPassword   string `mapstructure:"db-password"`
	DbSchema     string `mapstructure:"db-schema"`
}

func NewEnv() Config {

	// AddConfigPath called multiple times for testing purposes (viper look for config file from the path we call NewEnv)
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../../../")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	config := Config{}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("☠️ cannot read configuration", err)
		} else {
			fmt.Println("☠️ config file was found but another error was produced", err)
		}
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("☠️ environment can't be loaded: ", err)
	}

	ForceMapping(&config)

	return config
}

func ForceMapping(env *Config) {

	if env.Environment == "" {
		env.Environment = viper.GetString("environment")
	}

	if env.SmtpServer == "" {
		env.SmtpServer = viper.GetString("smtp-server")
	}

	if env.SmtpPort == "" {
		env.SmtpPort = viper.GetString("smtp-port")
	}

	if env.SmtpUsername == "" {
		env.SmtpUsername = viper.GetString("smtp-username")
	}

	if env.SmtpPassword == "" {
		env.SmtpPassword = viper.GetString("smtp-password")
	}

	if env.GeneralEmail == "" {
		env.GeneralEmail = viper.GetString("general-email")
	}

	if env.DbHost == "" {
		env.DbHost = viper.GetString("db-host")
	}

	if env.DbPort == "" {
		env.DbPort = viper.GetString("db-port")
	}

	if env.DbUsername == "" {
		env.DbUsername = viper.GetString("db-username")
	}

	if env.DbPassword == "" {
		env.DbPassword = viper.GetString("db-password")
	}

	if env.DbSchema == "" {
		env.DbSchema = viper.GetString("db-schema")
	}

}

var ConfigModule = fx.Provide(NewEnv)
