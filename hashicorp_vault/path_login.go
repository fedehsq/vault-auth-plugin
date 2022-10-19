package authplugin

import (
	"context"
	"time"
	"vault-auth-plugin/hashicorp_vault/api/user"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
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

	user, err := user.SignIn(username, password, b.jwt)
	if err != nil {
		return nil, err
	}

	// Compose the response
	resp := &logical.Response{
		Auth: &logical.Auth{
			InternalData: map[string]interface{}{
				"password": user.Password,
			},
			// Policies can be passed in as a parameter to the request
			Policies: []string{"ssh-policy"},
			Metadata: map[string]string{
				"username": user.Username,
			},
			// Lease options can be passed in as parameters to the request
			LeaseOptions: logical.LeaseOptions{
				TTL:       3600 * time.Second,
				MaxTTL:    7200 * time.Second,
				Renewable: true,
			},
		},
	}

	return resp, nil
}
