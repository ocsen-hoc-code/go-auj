package userrepository

import (
	"context"
	"database/sql"

	"github.com/ocsen-hoc-code/go-auj/models/config"
	"github.com/ocsen-hoc-code/go-auj/models/user"
)

type UserRepository struct {
	dbConfig *config.DbConfig
}

func NewUserRepository(dbInfo *config.DbConfig) *UserRepository {
	return &UserRepository{dbConfig: dbInfo}
}

// FindUser returns tasks if match userID AND password
func (rp UserRepository) FindUser(ctx context.Context, userName, password string) (*user.User, bool) {
	stmt := `SELECT id, username FROM users WHERE username = $1 AND password = $2`
	row := rp.dbConfig.Database.QueryRowContext(ctx, stmt,
		sql.NullString{String: userName, Valid: true},
		sql.NullString{String: password, Valid: true})

	user := &user.User{}
	err := row.Scan(&user.ID, &user.UserName)
	if err != nil {
		return user, false
	}

	return user, true
}
