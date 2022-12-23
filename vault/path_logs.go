package authplugin

import (
	"context"
	logapi "github.com/fedehsq/vault/api/log"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func (b *backend) pathLogs() *framework.Path {
	return &framework.Path{
		Pattern: "logs",
		// Add the query parameter
		Fields: map[string]*framework.FieldSchema{
			"q": {
				Type:        framework.TypeString,
				Description: "Query to filter the logs",
				Query:       true,
				Default:     "",
			},
			"command": {
				Type:        framework.TypeString,
				Description: "Command of the user in the remote machine",
				Query:       false,
			},
			"username": {
				Type:        framework.TypeString,
				Description: "Username of the user in the remote machine",
				Query:       false,
			},
			"ssh_address": {
				Type:        framework.TypeString,
				Description: "IP of the remote machine",
				Query:       false,
			},
		},

		Operations: map[logical.Operation]framework.OperationHandler{
			logical.ReadOperation: &framework.PathOperation{
				Callback: b.readLogs,
				Summary:  "Get the requested logs. If no parameters are passed, get all of them.",
			},
			logical.CreateOperation: &framework.PathOperation{
				Callback: b.createLog,
				Summary:  "Create a new log",
			},
			
			logical.UpdateOperation: &framework.PathOperation{
				Callback: b.createLog,
				Summary:  "Create a new log",
			},

		},
	}
}

func (b *backend) readLogs(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	// Get the parameters from command line
	q := data.Get("q").(string)
	// Get the JWT from the vault storage
	jwt, err := getJWT(ctx, req.Storage)
	if err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}
	logs, err := logapi.Get(jwt, q)
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

func (b *backend) createLog(ctx context.Context,
	req *logical.Request,
	data *framework.FieldData) (*logical.Response, error) {
	command := data.Get("command").(string)
	sshAddress := data.Get("ssh_address").(string)
	username := data.Get("username").(string)

	if sshAddress == "" || username == "" {
		return logical.ErrorResponse("ssh_address and username are required"), nil
	}

	// Get the JWT from the vault storage
	jwt, err := getJWT(ctx, req.Storage)
	if err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}
	err = logapi.Create(command, sshAddress, username, jwt)

	if err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}
	// Return the status code of the operation (201)
	return &logical.Response{
		Data: map[string]interface{}{
			"status": "201",
		},
	}, nil	
}
