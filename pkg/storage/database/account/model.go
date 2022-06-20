package account

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"iago-effting/api-example/pkg/accounts"
	"iago-effting/api-example/pkg/authentication"
	"iago-effting/api-example/pkg/storage/database"
)

type UserBun struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID       uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	Email    string
	Password string
}

func (u UserBun) ToEntity() accounts.User {
	return accounts.User{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
	}
}

type Repository struct {
	connection *bun.DB
}

func Repo() Repository {
	return Repository{
		connection: database.BunDb,
	}
}

func (r Repository) Get(ctx context.Context, id string) (*accounts.User, error) {
	var userModel UserBun
	var user accounts.User

	err := r.connection.NewSelect().Model(&userModel).Scan(ctx)
	if err != nil {
		return &user, err
	}

	user = userModel.ToEntity()

	return &user, nil
}

func (r Repository) Delete(ctx context.Context, id string) error {
	var user UserBun
	_, err := r.connection.NewDelete().Model(&user).Where("id = ?", id).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) Save(ctx context.Context, user accounts.User) (*accounts.User, error) {
	password, errHash := authentication.HashPassword(user.Password)
	if errHash != nil {
		return nil, errHash
	}

	model := &UserBun{
		Email:    user.Email,
		Password: password,
	}

	_, err := r.connection.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}

	user = model.ToEntity()

	return &user, nil
}

func (r Repository) All(ctx context.Context) (*[]accounts.User, error) {
	var usersModel []UserBun
	var users []accounts.User

	err := r.connection.NewSelect().Model(&usersModel).Scan(ctx)
	if err != nil {
		return &users, err
	}

	for _, user := range usersModel {
		users = append(users, user.ToEntity())
	}

	return &users, nil
}
