package authplugin

import (
	"context"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"time"
	"vault-auth-plugin/hashicorp_vault/api/user"
)

// Handle user login
func (b *backend) pathLogin() *framework.Path {
	return &framework.Path{
		Pattern: "login$",
		Fields: map[string]*framework.FieldSchema{
			"username": {
				Type:        framework.TypeString,
				Description: "Username of the user",
			},
			"password": {
				Type:        framework.TypeString,
				Description: "Password of the user",
			},
		},
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.UpdateOperation: &framework.PathOperation{
				Callback: b.handleLogin,
				Summary:  "Log in using a username and password",
			},
		},
	}
}

func (b *backend) handleLogin(ctx context.Context,
	req *logical.Request,
	data *framework.FieldData) (*logical.Response, error) {

	username := data.Get("username").(string)
	if username == "" {
		return logical.ErrorResponse("username must be provided"), nil
	}

	password := data.Get("password").(string)
	if password == "" {
		return logical.ErrorResponse("password must be provided"), nil
	}

	// Get the JWT from the vault storage
	JWT, err := getJWT(ctx, req.Storage)
	if err != nil {
		return nil, err
	}
	user, err := user.SignIn(username, password, JWT)
	if err != nil {
		return nil, err
	}

	// Compose the response
	resp := &logical.Response{
		Auth: &logical.Auth{
			InternalData: map[string]interface{}{
				"password": user.Password,
			},
			Policies: []string{"ssh-policy"},
			Metadata: map[string]string{
				"username": user.Username,
			},
			LeaseOptions: logical.LeaseOptions{
				TTL:       60 * time.Minute,
				MaxTTL:    60 * time.Minute,
				Renewable: false,
			},
		},
	}

	return resp, nil
}
