package repo

import (
	"context"
	"github.com/LeoUraltsev/taskchecker/internal/models"
	"github.com/LeoUraltsev/taskchecker/internal/repo/pgdb"
	"github.com/LeoUraltsev/taskchecker/pkg/postgres"
)

type User interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	GetUserById(userID int) (models.User, error)
	GetUserByUsernameAndPassword(username string, password string) (models.User, error)
}

type Repository struct {
	User
}

func NewRepository(pg *postgres.Postgres) *Repository {
	return &Repository{User: pgdb.NewUserRepo(pg)}
}
