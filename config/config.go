package config

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	ConnURI  string
	Database string
}

func init() {
	viper.SetDefault("API_PORT", "9000")
	viper.SetDefault("DATABASE_URL", "mongodb://localhost:27017")
}

func LoadConfig() error {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("API_PORT"),
	}

	cfg.DB = DBConfig{
		ConnURI:  viper.GetString("DATABASE_URL"),
		Database: viper.GetString("DATABASE_NAME"),
	}

	return nil
}

func GetDBConfig() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
