package authplugin

import (
	"context"
	userapi "vault-auth-plugin/vault/api/user"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func (b *backend) pathUser() []*framework.Path {
	return []*framework.Path{
		{
			Pattern: "user/" + framework.GenericNameRegex("username"),

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
	JWT, err := getJWT(ctx, req.Storage)
	if err != nil {
		return nil, err
	}
	// check if the user already exists
	u, _ := userapi.GetUser(username, JWT)
	if u != nil {
		// Update the user
		_, err := userapi.UpdateUser(username, password, JWT)
		if err != nil {
			return logical.ErrorResponse(err.Error()), nil
		}
	} else {
		// Store to db
		_, err := userapi.SignUp(username, password, JWT)
		if err != nil {
			return logical.ErrorResponse(err.Error()), nil
		}
	}
	return nil, nil
}

func (b *backend) handleUserDelete(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	username := data.Get("username").(string)
	if username == "" {
		return logical.ErrorResponse("username must be provided"), nil
	}

	// Get the JWT from the vault storage
	JWT, err := getJWT(ctx, req.Storage)
	if err != nil {
		return nil, err
	}
	err = userapi.DeleteUser(username, JWT)
	if err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}

	return nil, nil
}

// func (b *backend) handleExistenceCheck(ctx context.Context, req *logical.Request, data *framework.FieldData) (bool, error) {
// 	username := data.Get("username").(string)
// 	if username == "" {
// 		return false, nil
// 	}
// 	_, err := user.GetUser(username, b.jwt)
// 	if err != nil {
// 		return false, errors.New("failed to get user")
// 	}
// 	return true, nil
// }
