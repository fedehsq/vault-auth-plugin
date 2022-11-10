package authplugin

import (
	"context"
	"github.com/fedehsq/vault/api/user"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func (b *backend) pathUsers() *framework.Path {
	return &framework.Path{
		Pattern: "user/get-all/?$",
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.ListOperation: &framework.PathOperation{
				Callback: b.handleUsers,
				Summary:  "List existing users.",
			},
		},
	}
}

func (b *backend) handleUsers(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	// Get the JWT from the vault storage
	JWT, err := getJWT(ctx, req.Storage)
	if err != nil {
		return nil, err
	}
	users, err := userapi.GetUsers(JWT)
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
