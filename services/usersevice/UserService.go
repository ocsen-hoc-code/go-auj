package userservice

import (
	"context"

	"github.com/ocsen-hoc-code/go-auj/models/config"
	"github.com/ocsen-hoc-code/go-auj/models/user"
	userrepository "github.com/ocsen-hoc-code/go-auj/reponsitories/user"
	"github.com/ocsen-hoc-code/go-auj/utils/jwtutil"
)

type UserService struct {
	reps      *userrepository.UserRepository
	jwtConfig *config.JWTConfig
}

func NewUserService(userRepsInject *userrepository.UserRepository, jwtConfigInject *config.JWTConfig) *UserService {
	return &UserService{reps: userRepsInject, jwtConfig: jwtConfigInject}
}

func (serv *UserService) Login(ctx context.Context, user user.User) (string, string) {
	msgError := ""
	token := ""

	u, ok := serv.reps.FindUser(ctx, user.UserName, user.Password)

	if !ok {
		msgError = "incorrect user_id/pwd"
	}

	token, err := jwtutil.CreateToken(*u, serv.jwtConfig.SecretKey, serv.jwtConfig.ExpireTime)
	if err != nil {
		msgError = err.Error()
	}

	return token, msgError
}
