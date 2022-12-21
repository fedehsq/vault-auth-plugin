path "auth/auth-plugin/users/*" {
    capabilities = ["create", "update", "delete", "read"]
}

path "auth/auth-plugin/logs" {
    capabilities = ["read", "create", "update", "delete"]
}

path "secret/data/bastion" {
    capabilities = ["read"]
}