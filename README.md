# Custom Auth plugin for [HashiCorp Vault](https://www.vaultproject.io/)

## Instr
go build -o hashicorp_vault/plugins/auth-plugin hashicorp_vault/cmd/main.go
vault server -dev -dev-root-token-id=root -dev-plugin-dir=./hashicorp_vault/plugins
export VAULT_ADDR="http://127.0.0.1:8200"
vault auth enable -path=auth-plugin auth-plugin
vault write auth/auth-plugin/admin-login username=admin password=admin
vault write auth/auth-plugin/login username=fede password=pwd
