package config

import "github.com/spf13/viper"

type Config struct {
	VaultAddress         string `mapstructure:"VAULT_ADDRESS"`
	VaultServerAddress   string `mapstructure:"VAULT_SERVER_ADDRESS"`
	VaultServerDbAddress string `mapstructure:"VAULT_SERVER_DB_ADDRESS"`
	VaultServerDbPort    int    `mapstructure:"VAULT_SERVER_DB_PORT"`
	VaultServerDbUser    string `mapstructure:"VAULT_SERVER_DB_USER"`
	VaultServerDbName    string `mapstructure:"VAULT_SERVER_DB_NAME"`
	BastionHostAddress   string `mapstructure:"BASTION_HOST_ADDRESS"`
	SshHost              string `mapstructure:"SSH_HOST"`
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
