package configs

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
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	viper.SetDefault("api.Port", "9000")
	viper.SetDefault("datbase.Host", "localhost")
	viper.SetDefault("datbase.Port", "5432")
}

func load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.Port"),
	}
	CFG.DB = DBConfig{
		Host:     viper.GetString("database.Host"),
		Port:     viper.GetString("database.Port"),
		User:     viper.GetString("database.User"),
		Pass:     viper.GetString("database.Pass"),
		Database: viper.GetString("database.Database"),
	}
	return nil
}

func GetDB() *DBconfig {
	return cfg.DB
}
