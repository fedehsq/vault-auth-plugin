package authplugin

import (
	"context"
	"time"

	"github.com/fedehsq/vault/api/user"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

// Handle user login
func (b *backend) pathLogin() *framework.Path {
	return &framework.Path{
		Pattern: "users/signin$",
		Fields: map[string]*framework.FieldSchema{
			"username": {
				Type:        framework.TypeString,
				Description: "Username of the user",
			},
			"password": {
				Type:        framework.TypeString,
				Description: "Password of the user",
			},
			"jwt": {
				Type:        framework.TypeString,
				Description: "JWT to be used for authentication",
			},
		},
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.UpdateOperation: &framework.PathOperation{
				Callback: b.handleLogin,
				Summary:  "Log in using a username and password.",
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

	jwt := data.Get("jwt").(string)
	if jwt == "" {
		return logical.ErrorResponse("jwt must be provided"), nil
	}

	user, err := userapi.SignIn(username, password, jwt)
	if err != nil {
		return nil, err
	}

	// Compose the response
	resp := &logical.Response{
		Auth: &logical.Auth{
			InternalData: map[string]interface{}{
				"password": user.Password,
			},
			Policies: []string{"user-policy"},
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
