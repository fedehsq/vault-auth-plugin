package authplugin

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

// backend wraps the backend framework and adds a map for storing key value pairs.
type backend struct {
	*framework.Backend
	jwt string
}

var _ logical.Factory = Factory

// Factory configures and returns Mock backends
func Factory(ctx context.Context, conf *logical.BackendConfig) (logical.Backend, error) {
	b, err := newBackend()
	if err != nil {
		return nil, err
	}

	if conf == nil {
		return nil, fmt.Errorf("configuration passed into backend is nil")
	}

	if err := b.Setup(ctx, conf); err != nil {
		return nil, err
	}

	return b, nil
}

func newBackend() (*backend, error) {
	b := new(backend)

	b.Backend = &framework.Backend{
		Help:        strings.TrimSpace(mockHelp),
		BackendType: logical.TypeCredential,
		PathsSpecial: &logical.Paths{
			Unauthenticated: []string{
				"admin-login",
				"login",
			},
		},
		// AuthRenew:   b.adminAuthRenew,
		Paths: framework.PathAppend(
			[]*framework.Path{
				b.adminPathLogin(),
				b.pathLogin(),
				b.pathUsers(),
			},
			b.pathUser(),
		),
	}

	return b, nil
}

const mockHelp = `
The backend is a auth backend that can manage user stores user and password data in a remote vault database.
Only the bastion host is allowed to login to the backend and to manage users.
Each request to the API of the remote server is authenticated with a JWT token that is generated by the remote server when the admin logs in.
When an user logs in, the backend will check if the user exists in the database and if the password is correct. 
In case of success, the vault backend will return a token that can be used.
`
