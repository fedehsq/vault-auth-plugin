package authplugin

import (
	"context"

	"github.com/fedehsq/vault/api/user"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func (b *backend) pathUser() []*framework.Path {
	return []*framework.Path{
		{
			Pattern: "users/" + framework.GenericNameRegex("username"),

			Fields: map[string]*framework.FieldSchema{
				"username": {
					Type:        framework.TypeString,
					Description: "Specifies the username",
				},
				"password": {
					Type:        framework.TypeString,
					Description: "Specifies the password for the user",
				},
			},

			Operations: map[logical.Operation]framework.OperationHandler{
				logical.UpdateOperation: &framework.PathOperation{
					Callback: b.handleUserWrite,
					Summary:  "Adds a new user to the auth method.",
				},
				logical.CreateOperation: &framework.PathOperation{
					Callback: b.handleUserWrite,
					Summary:  "Updates a user on the auth method.",
				},
				logical.DeleteOperation: &framework.PathOperation{
					Callback: b.handleUserDelete,
					Summary:  "Deletes a user on the auth method.",
				},
				logical.ReadOperation: &framework.PathOperation{
					Callback: b.handleUserRead,
					Summary:  "Reads a user on the auth method.",
				},
			},
		},
	}
}

func (b *backend) handleUserWrite(ctx context.Context,
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
	jwt, err := getJWT(ctx, req.Storage)
	if err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}
	// check if the user already exists
	u, _ := userapi.GetUser(username, jwt)
	if u != nil {
		// Update the user
		_, err := userapi.UpdateUser(username, password, jwt)
		if err != nil {
			return logical.ErrorResponse(err.Error()), nil
		}
	} else {
		// Store to db
		_, err := userapi.SignUp(username, password, jwt)
		if err != nil {
			return logical.ErrorResponse(err.Error()), nil
		}
	}
	return nil, nil
}

func (b *backend) handleUserRead(ctx context.Context,
	req *logical.Request,
	data *framework.FieldData) (*logical.Response, error) {
	username := data.Get("username").(string)
	if username == "" {
		return logical.ErrorResponse("username must be provided"), nil
	}

	// Get the JWT from the vault storage
	jwt, err := getJWT(ctx, req.Storage)
	if err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}
	u, err := userapi.GetUser(username, jwt)
	if err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}
	if u == nil {
		return logical.ErrorResponse("user not found"), nil
	}
	return &logical.Response{
		Data: map[string]interface{}{
			"username": u.Username,
		},
	}, nil
}

func (b *backend) handleUserDelete(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	username := data.Get("username").(string)
	if username == "" {
		return logical.ErrorResponse("username must be provided"), nil
	}
	// Get the JWT from the vault storage
	jwt, err := getJWT(ctx, req.Storage)
	if err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}
	err = userapi.DeleteUser(username, jwt)
	if err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}

	return nil, nil
}
