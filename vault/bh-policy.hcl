path "auth/auth-plugin/users/*" {
    capabilities = ["list"]
}

path "auth/auth-plugin/user/*" {
    capabilities = ["create", "update", "delete"]
}

path "auth/auth-plugin/logs/*" {
    capabilities = ["list"]
}