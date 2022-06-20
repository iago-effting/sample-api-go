package auth

import "context"

type Repository interface {
	Save(ctx context.Context, user User) (*User, error)
}
