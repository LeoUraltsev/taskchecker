package pgdb

import (
	"context"
	"github.com/LeoUraltsev/taskchecker/internal/models"
	"github.com/LeoUraltsev/taskchecker/pkg/postgres"
)

type UserRepo struct {
	*postgres.Postgres
}

func (u *UserRepo) CreateUser(ctx context.Context, user models.User) (int, error) {
	q := ``
	var id int
	if err := u.Pool.QueryRow(ctx, q).Scan(&id); err != nil {

	}
	return 0, nil
}

func (u *UserRepo) GetUserById(userID int) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) GetUserByUsernameAndPassword(username string, password string) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}
