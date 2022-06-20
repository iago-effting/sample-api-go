package accounts

import (
	"context"
)

type Repository interface {
	Save(ctx context.Context, user User) (*User, error)
	All(ctx context.Context) (*[]User, error)
	Get(ctx context.Context, id string) (*User, error)
	GetBy(ctx context.Context, field string, value string) (*User, error)
	Delete(ctx context.Context, id string) error
}
