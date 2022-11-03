package main

import (
	pluginBackend "github.com/fedehsq/vault-auth-plugin/vault"
	"github.com/fedehsq/vault-auth-plugin/vault/config"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/sdk/plugin"
	"log"
	"os"
)

// Use the Vault SDK's plugin library to start the plugin
// and communicate with the Vault API.
func main() {
	err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	apiClientMeta := &api.PluginAPIClientMeta{}
	flags := apiClientMeta.FlagSet()
	flags.Parse(os.Args[1:])

	tlsConfig := apiClientMeta.GetTLSConfig()
	tlsProviderFunc := api.VaultPluginTLSProvider(tlsConfig)

	err = plugin.Serve(&plugin.ServeOpts{
		BackendFactoryFunc: pluginBackend.Factory,
		TLSProviderFunc:    tlsProviderFunc,
	})
	if err != nil {
		logger := hclog.New(&hclog.LoggerOptions{})

		logger.Error("plugin shutting down", "error", err)
		os.Exit(1)
	}
}
