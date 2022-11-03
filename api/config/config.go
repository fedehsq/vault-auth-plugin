package config

import "github.com/spf13/viper"

type Config struct {
	VaultAddress  string `mapstructure:"VAULT_ADDRESS"`
	ApiAddress    string `mapstructure:"API_ADDRESS"`
	DbAddress     string `mapstructure:"DB_ADDRESS"`
	DbPort        int    `mapstructure:"DB_PORT"`
	DbUser        string `mapstructure:"DB_USER"`
	DbName        string `mapstructure:"DB_NAME"`
	ApiVaultToken string `mapstructure:"API_VAULT_TOKEN"`
}

var Conf *Config

func LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		return err
	}
	return nil
}
