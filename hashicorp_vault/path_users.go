package authplugin

import (
	"context"
	"crypto/subtle"
	"errors"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"time"
	"vault-auth-plugin/hashicorp_vault/api/user"
)

func (b *backend) pathUsers() []*framework.Path {
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
				logical.ListOperation: &framework.PathOperation{
					Callback: b.handleUsersList,
					Summary:  "List existing users.",
				},
			},
			ExistenceCheck: b.handleExistenceCheck,
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

	// Store to db
	_, err := user.SignUp(username, password, b.jwt)
	if err != nil {
		return logical.ErrorResponse("failed to create client"), nil
	}

	return nil, nil
}

func (b *backend) handleUserDelete(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	username := data.Get("username").(string)
	if username == "" {
		return logical.ErrorResponse("username must be provided"), nil
	}

	if user.DeleteUser(username, b.jwt) != nil {
		return logical.ErrorResponse("failed to delete user"), nil
	}

	return nil, nil
}

func (b *backend) pathAuthRenew(ctx context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {
	username := req.Auth.Metadata["username"]
	pw := req.Auth.InternalData["password"].(string)

	user, err := user.GetUser(username, b.jwt)
	if err != nil {
		return nil, err
	}

	if subtle.ConstantTimeCompare([]byte(pw), []byte(user.Password)) != 1 {
		return nil, errors.New("internal data does not match")
	}

	resp := &logical.Response{Auth: req.Auth}
	resp.Auth.TTL = 30 * time.Second
	resp.Auth.MaxTTL = 60 * time.Minute

	return resp, nil
}

func (b *backend) handleExistenceCheck(ctx context.Context, req *logical.Request, data *framework.FieldData) (bool, error) {
	username := data.Get("name").(string)
	_, err := user.GetUser(username, b.jwt)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (b *backend) handleUsersList(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	users, err := user.GetUsers(b.jwt)
	if err != nil {
		return nil, err
	}
	// Iterate over the users and add them to the list
	userNames := make([]string, 0, len(users))
	for _, user := range users {
		userNames = append(userNames, user.Username)
	}

	return logical.ListResponse(userNames), nil
}
