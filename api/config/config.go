package config

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/spf13/viper"
)

type Config struct {
	VaultAddress     string `mapstructure:"VAULT_ADDRESS"`
	ApiAddress       string `mapstructure:"API_ADDRESS"`
	ElasticSearch    string `mapstructure:"ELASTICSEARCH_URL"`
	ElasticSearchPwd string `mapstructure:"ELASTIC_PASSWORD"`
	DbAddress        string `mapstructure:"DB_ADDRESS"`
	DbPort           int    `mapstructure:"DB_PORT"`
	DbUser           string `mapstructure:"DB_USER"`
	DbName           string `mapstructure:"DB_NAME"`
	DbPassword       string `mapstructure:"DB_PASSWORD"`
	ApiVaultToken    string `mapstructure:"VAULT_TOKEN"`
	Develop          int    `mapstructure:"DEVELOP"`
}

var Conf *Config
var EsClient *elasticsearch.Client

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
	// Set the Elasticsearch client with url from config
	EsClient, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			Conf.ElasticSearch,
		},
		Username: "elastic",
		Password: Conf.ElasticSearchPwd,
	})
	if err != nil {
		return err
	}
	return nil
}
