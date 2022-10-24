# Implementation of a controlled access system using Bastion Host and Vault
Increasing automation in IT processes and the evolution of software lifecycle processes from a DevOps perspective have meant that there is less and less need to perform direct access to IT systems: typically, access is for the purpose of performing critical operations and extraordinary maintenance.  
Access to such systems therefore should be as limited as possible and subject to strict monitoring of the accesses performed.  
To this end, it is useful to define a single point of access on which to focus security audits: such a system is typically referred to as a bastion host.  
Another critical point in system access is the protection of access keys: these in the enterprise environment are managed in a wide variety of ways, sometimes without any use of protective measures, others with increasingly sophisticated systems such as vaults or HSMs.

The purpose of this thesis project is to implement a system based on bastion hosts, which through the implementation of an authorization workflow, allows granting or denying access to remote systems through the automatic use of keys retrieved from a vault by a bastion host.

## Auth plugin for [HashiCorp Vault](https://www.vaultproject.io/)
The first step in the implementation of the system is the development of an authentication plugin for Vault.  
The plugin is based on the [plugin development guide](https://www.vaultproject.io/docs/internals/plugins.html) provided by HashiCorp and is written in Go.  
The workflow of the plugin under development is as follows:
```mermaid
sequenceDiagram
    actor User
    actor Operator
    participant Bastion Host
    participant Vault
    participant Vault Server
    participant Target Host
    note over Operator: Vault setup: enable and write plugin policies using root token
    Operator->>Vault: Vault setup
    User->>Bastion Host: Authentication over Bastion Host
    Bastion Host->>Vault: Bastion Host authentication
    Vault->>Vault Server: Bastion Host Credentials
    Note over Vault Server: JWT creation to call the other API
    Vault Server->>Vault: JWT 
    note over Vault: Vault Token creation with the plug-in policies
    Vault->>Bastion Host: Vault Token 
    note over Bastion Host: Forwards the user credentials to the Vault Server using the Vault Token and JWT
    Bastion Host->>Vault: User Credentials
    note over Vault:Vault Token checks
    Vault->>Vault Server: User Credentials
    note over Vault Server: JWT checks
    Vault Server->>Vault: OK
    note over Vault: Detaches a valid token for the User
    Vault->>Bastion Host: Vault Token
    note over Bastion Host: Uses the user Vault token to request the OTP
    Bastion Host->>Vault: Get OTP
    note over Vault:Vault User Token checks
    Vault->>Bastion Host: OTP
    note over Bastion Host: !!! THE FOLLOWING STEPS MUST BE AUTOMATIZED !!!
    Bastion Host->>User: OTP
    User->>Target Host: ssh address OTP

```

## Instructions
1. Edit the env file with your own values

    ```
    nano .env
    ```

2. Setup the remote host following the [hashicorp guide for SSH](https://learn.hashicorp.com/tutorials/vault/ssh-otp?in=vault/secrets-management).  
If you have already followed the above instructions before, do this:
    - Starts vagrant
        ```
        vagrant up
        ```

    - Connect to the remote host
        ```
        vagrant ssh
        ```

    - Change the 'vault_addr' variable with the address of your vault server
        ```
        sudo nano /etc/vault-ssh-helper.d/config.hcl
        ```

    - Restart the service
        ```
        sudo systemctl restart sshd
        ```

    - Verify that the configuration is correct

        ```
        vault-ssh-helper -verify-only -dev -config /etc/vault-ssh-helper.d/config.hcl
        ```
        
    - Exit from the remote host
        ```
        exit
        ```

3. Starts the vault server
    ```
    go run vault_server/cmd/main.go
    ```

4. Starts the bastion host
    ```
    go run bastion_host/cmd/main.go
    ```

5. Build and starts the vault plugin
    ```
    make
    ```

6. Setup the vault and the plugin
    ```
    export VAULT_ADDR=http://$(ipconfig getifaddr en0):8200
    make vault-setup
    ```