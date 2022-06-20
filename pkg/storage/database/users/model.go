package users

import (
	"context"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"iago-effting/api-example/pkg/auth"
	"iago-effting/api-example/pkg/storage/database"
)

type UserBun struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID       uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	Email    string
	Password string
}

func (u UserBun) ToEntity() auth.User {
	return auth.User{
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

func (r Repository) Get(ctx context.Context, id string) (*auth.User, error) {
	var userModel UserBun
	var user auth.User

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

func (r Repository) Save(ctx context.Context, user auth.User) (*auth.User, error) {
	model := &UserBun{
		Email:    user.Email,
		Password: user.Password,
	}

	_, err := r.connection.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}

	user = model.ToEntity()

	return &user, nil
}

func (r Repository) All(ctx context.Context) (*[]auth.User, error) {
	var usersModel []UserBun
	var users []auth.User

	err := r.connection.NewSelect().Model(&usersModel).Scan(ctx)
	if err != nil {
		return &users, err
	}

	for _, user := range usersModel {
		users = append(users, user.ToEntity())
	}

	return &users, nil
}