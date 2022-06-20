package authentication

import "context"

type Repository interface {
	CheckCredentials(ctx context.Context, credentials Credentials) (bool, error)
}
