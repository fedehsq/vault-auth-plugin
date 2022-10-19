package authplugin

import (
	"context"
	"time"
	"vault-auth-plugin/hashicorp_vault/api/admin"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

// Handle admin login
func (b *backend) adminPathLogin() *framework.Path {
	return &framework.Path{
		Pattern: "admin-login$",
		Fields: map[string]*framework.FieldSchema{
			"username": {
				Type:        framework.TypeString,
				Description: "Username of the admin",
			},
			"password": {
				Type:        framework.TypeString,
				Description: "Password of the admin",
			},
		},
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.UpdateOperation: &framework.PathOperation{
				Callback: b.handleAdminLogin,
				Summary:  "Log in using a username and password",
			},
		},
	}
}

func (b *backend) handleAdminLogin(ctx context.Context,
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

	admin, err := admin.SignIn(username, password)
	if err != nil {
		return nil, err
	}

	b.jwt = admin.JWT
	
	// Compose the response
	resp := &logical.Response{
		Auth: &logical.Auth{
			InternalData: map[string]interface{}{
				"password": admin.Password,
			},
			// Policies can be passed in as a parameter to the request
			Policies: []string{"plugin-policy"},
			Metadata: map[string]string{
				"username": admin.Username,
			},
			LeaseOptions: logical.LeaseOptions{
				TTL:       3600 * time.Second,
				MaxTTL:    7200 * time.Second,
				Renewable: true,
			},
		},
	}

	return resp, nil
}
