package authplugin

import (
	"context"

	"github.com/fedehsq/vault/api/remote_host_users"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func (b *backend) remoteHostUsersPath() *framework.Path {
	return &framework.Path{
		Pattern: "remote-host-users$",
		Fields: map[string]*framework.FieldSchema{
			"ip": {
				Type:        framework.TypeString,
				Description: "Ip of the remote host",
			},
			"username": {
				Type:        framework.TypeString,
				Description: "Username to be added to the remote host",
			},
		},
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.UpdateOperation: &framework.PathOperation{
				Callback: b.createRemoteHostUser,
				Summary:  "Update an authorized user for a remote host.",
			},
			logical.CreateOperation: &framework.PathOperation{
				Callback: b.createRemoteHostUser,
				Summary:  "Create an authorized user for a remote host.",
			},
			logical.DeleteOperation: &framework.PathOperation{
				Callback: b.deleteRemoteHostUser,
				Summary:  "Delete an authorized user for a remote host.",
			},
			logical.ReadOperation: &framework.PathOperation{
				Callback: b.getRemoteHostUser,
				Summary:  "Read an authorized user for a remote host or all of them.",
			},
		},
	}
}

func (b *backend) getRemoteHostUser(ctx context.Context,
	req *logical.Request,
	data *framework.FieldData) (*logical.Response, error) {

	ip := data.Get("ip").(string)
	username := data.Get("username").(string)
	if ip == "" {
		return logical.ErrorResponse("ip must be provided"), nil
	}
	// Get the JWT from the vault storage
	jwt, err := getJWT(ctx, req.Storage)
	if err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}
	if username == "" {
		remoteHostUsers, err := remotehostusersapi.GetAll(ip, jwt)
		if err != nil {
			return nil, err
		}
		resp := &logical.Response{
			Data: remoteHostUsers.ToMap(),
		}
		return resp, nil
	} else {
		remoteHostUser, err := remotehostusersapi.Get(ip, username, jwt)
		if err != nil {
			return nil, err
		}
		resp := &logical.Response{
			Data: remoteHostUser.ToMap(),
		}
		return resp, nil
	}
}

func (b *backend) deleteRemoteHostUser(ctx context.Context,
	req *logical.Request,
	data *framework.FieldData) (*logical.Response, error) {

	ip := data.Get("ip").(string)
	username := data.Get("username").(string)
	if ip == "" {
		return logical.ErrorResponse("ip must be provided"), nil
	}
	if username == "" {
		return logical.ErrorResponse("username must be provided"), nil
	}

	jwt, err := getJWT(ctx, req.Storage)
	if err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}
	err = remotehostusersapi.Delete(ip, username, jwt)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (b *backend) createRemoteHostUser(ctx context.Context,
	req *logical.Request,
	data *framework.FieldData) (*logical.Response, error) {

	ip := data.Get("ip").(string)
	username := data.Get("username").(string)
	if ip == "" {
		return logical.ErrorResponse("ip must be provided"), nil
	}
	if username == "" {
		return logical.ErrorResponse("username must be provided"), nil
	}

	jwt, err := getJWT(ctx, req.Storage)
	if err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}
	_, err = remotehostusersapi.Create(ip, username, jwt)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
