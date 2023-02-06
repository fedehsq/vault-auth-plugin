path "auth/auth-plugin/users/*" {
    capabilities = ["create", "update", "delete", "read"]
}

path "auth/auth-plugin/logs" {
    capabilities = ["read", "create", "update", "delete"]
}

path "secret/data/bastion" {
    capabilities = ["read"]
}

path "auth/auth-plugin/remote-hosts" {
    capabilities = ["read", "create", "update", "delete"]
}

path "auth/auth-plugin/remote-host-users" {
    capabilities = ["read", "create", "update", "delete"]
}