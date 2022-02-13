package config

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		Postgres PostgresConfig
		GRPC     GRPCConfig
		Services map[string]string
	}

	PostgresConfig struct {
		Host     string
		Port     string
		User     string
		DBName   string
		Password string
	}

	GRPCConfig struct {
		Port string
	}
)

func Init(configsDir string) (*Config, error) {
	if err := parseConfigFile(configsDir); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("postgres", &cfg.Postgres); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("services", &cfg.Services); err != nil {
		return err
	}

	return viper.UnmarshalKey("grpc", &cfg.GRPC)
}

func parseConfigFile(folder string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.MergeInConfig()
}
