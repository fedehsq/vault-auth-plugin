package authplugin

import (
	"context"

	logapi "github.com/fedehsq/vault/api/log"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func (b *backend) pathLogs() *framework.Path {
	return &framework.Path{
		Pattern: "logs/?$",
		// Add the query parameter
		Fields: map[string]*framework.FieldSchema{
			"q": {
				Type:        framework.TypeString,
				Description: "Query to filter the logs",
				Query:       true,
				Default:     "",
			},
		},
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.ReadOperation: &framework.PathOperation{
				Callback: b.handleLogs,
				Summary:  "Get the requested logs. If no parameters are passed, get all of them.",
			},
		},
	}
}

func (b *backend) handleLogs(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	// Get the parameters from command line
	q := data.Get("q").(string)
	// Get the JWT from the vault storage
	JWT, err := getJWT(ctx, req.Storage)
	if err != nil {
		return nil, err
	}
	logs, err := logapi.Get(JWT, q)
	if err != nil {
		return nil, err
	}
	// Iterate over the logs and add them to the list
	logNames := make([]string, 0, len(logs))
	for _, log := range logs {
		logNames = append(logNames, log.String())
	}

	return &logical.Response{
		Data: map[string]interface{}{
			"logs": logNames,
		},
	}, nil
}
