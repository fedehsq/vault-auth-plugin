#!/bin/bash
VAULT_RETRIES=5
echo "Vault is starting..."
until vault status > /dev/null 2>&1 || [ "$VAULT_RETRIES" -eq 0 ]; do
        echo "Waiting for vault to start...: $((VAULT_RETRIES--))"
        sleep 1
done
echo "Authenticating to vault..."
vault login token=root
echo "Initializing plugin..."
vault auth enable -path=auth-plugin auth-plugin
echo "Setting up bastion host credentials..."
vault kv put -mount=secret bastion username=admin password=admin
echo "Setting up API secret key..."
vault kv put -mount=secret api key=SUPERSECRETKEY
echo "Enabling ssh..."
vault secrets enable ssh
vault write ssh/roles/otp_key_role \
key_type=otp \
default_user=vagrant \
cidr_list=0.0.0.0/0
echo "Writing bastion host policies..."
vault policy write bh-policy ./bh-policy.hcl
echo "Writing user policies..."
vault policy write user-policy ./user-policy.hcl
echo "Writing API policies..."
vault policy write api-policy ./api-policy.hcl
echo "Creating bastion host token..."
vault token create -policy=bh-policy -id=CAESIFf-ixZPKDzG3_rYR8TcfveN-AfG_JSWJKz4itilwfTjGh4KHGh2cy5ZajltYTVwSTlLUXNZWDhjRERjYjRkUHg
echo "Creating API token..."
vault token create -policy=api-policy -id=BXESIFf-ixZPKDzG3_rYR8TcfveN-AfG_TSWJKz4itilwfTjph4KHGh2cy5ZajltYTVwSTlLUXNZWDhjRERjYjRkUHg
echo "Complete..."