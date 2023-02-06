package authplugin

import (
	"context"

	"github.com/fedehsq/vault/api/remote_host"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func (b *backend) remoteHostPath() *framework.Path {
	return &framework.Path{
		Pattern: "remote-hosts$",
		Fields: map[string]*framework.FieldSchema{
			"ip": {
				Type:        framework.TypeString,
				Description: "Ip of the remote host",
			},
		},
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.UpdateOperation: &framework.PathOperation{
				Callback: b.createRemoteHost,
				Summary:  "Update a remote host.",
			},
			logical.CreateOperation: &framework.PathOperation{
				Callback: b.createRemoteHost,
				Summary:  "Create a remote host.",
			},
			logical.DeleteOperation: &framework.PathOperation{
				Callback: b.deleteRemoteHost,
				Summary:  "Delete a remote host.",
			},
			logical.ReadOperation: &framework.PathOperation{
				Callback: b.getRemoteHost,
				Summary:  "Read a remote host.",
			},
		},
	}
}

func (b *backend) getRemoteHost(ctx context.Context,
	req *logical.Request,
	data *framework.FieldData) (*logical.Response, error) {

	ip := data.Get("ip").(string)
	if ip == "" {
		return logical.ErrorResponse("ip must be provided"), nil
	}
	// Get the JWT from the vault storage
	jwt, err := getJWT(ctx, req.Storage)
	if err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}
	remoteHost, err := remotehostapi.Get(ip, jwt)
	if err != nil {
		return nil, err
	}

	resp := &logical.Response{
		Data: map[string]interface{}{
			"ip": remoteHost.Ip,
		},
	}
	return resp, nil
}

func (b *backend) deleteRemoteHost(ctx context.Context,
	req *logical.Request,
	data *framework.FieldData) (*logical.Response, error) {

	ip := data.Get("ip").(string)
	if ip == "" {
		return logical.ErrorResponse("ip must be provided"), nil
	}
	jwt, err := getJWT(ctx, req.Storage)
	if err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}
	err = remotehostapi.Delete(ip, jwt)
	if err != nil {
		return nil, err
	}

	return nil, nil
}


func (b *backend) createRemoteHost(ctx context.Context,
	req *logical.Request,
	data *framework.FieldData) (*logical.Response, error) {

	ip := data.Get("ip").(string)
	if ip == "" {
		return logical.ErrorResponse("ip must be provided"), nil
	}

	jwt, err := getJWT(ctx, req.Storage)
	if err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}
	_, err = remotehostapi.Create(ip, jwt)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
