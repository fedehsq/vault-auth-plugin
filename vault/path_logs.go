package authplugin

import (
	"context"
	"github.com/fedehsq/vault/api/log"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func (b *backend) pathLogs() *framework.Path {
	return &framework.Path{
		Pattern: "logs/?$",
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.ListOperation: &framework.PathOperation{
				Callback: b.handleLogs,
				Summary:  "List existing logs.",
			},
		},
	}
}

func (b *backend) handleLogs(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	// Get the JWT from the vault storage
	JWT, err := getJWT(ctx, req.Storage)
	if err != nil {
		return nil, err
	}
	logs, err := logapi.GetAll(JWT)
	if err != nil {
		return nil, err
	}
	// Iterate over the logs and add them to the list
	logNames := make([]string, 0, len(logs))
	for _, log := range logs {
		logNames = append(logNames, log.String())
	}

	return logical.ListResponse(logNames), nil
}
