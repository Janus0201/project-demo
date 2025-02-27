package service

import (
	"context"
	"errors"

	"github.com/MrLittle05/Gomall/app/user/biz/dal/mysql"
	"github.com/MrLittle05/Gomall/app/user/biz/model"
	user "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	email := req.Email
	password := req.Password
	if email == "" || password == "" {
		return nil, errors.New("email or password is empty")
	}
	row, err := model.GetByEmail(mysql.DB, email)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, errors.New("user not found")
	}
	// password check
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(password))
	if err != nil {
		return nil, errors.New("password not match")
	}

	resp = &user.LoginResp{
		UserId: int32(row.ID),
	}
	return resp, nil
}
