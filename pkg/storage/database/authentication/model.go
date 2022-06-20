package authentication

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"iago-effting/api-example/pkg/authentication"
	"iago-effting/api-example/pkg/storage/database"
)

type UserBun struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID       uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	Email    string
	Password string
}

type Repository struct {
	connection *bun.DB
}

func Repo() Repository {
	return Repository{
		connection: database.BunDb,
	}
}

func (r Repository) CheckCredentials(ctx context.Context, credentials authentication.Credentials) (bool, error) {
	var userModel UserBun

	err := r.connection.
		NewSelect().
		Model(&userModel).
		Where("email = ?", credentials.Email).
		Scan(ctx)

	if err != nil {
		return false, err
	}

	if !authentication.CheckPasswordHash(credentials.Password, userModel.Password) {
		return false, nil
	}

	return true, nil
}
