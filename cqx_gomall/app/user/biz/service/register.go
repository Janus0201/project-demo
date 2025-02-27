package service

import (
	"context"
	"errors"

	"github.com/MrLittle05/Gomall/app/user/biz/dal/mysql"
	"github.com/MrLittle05/Gomall/app/user/biz/model"
	user "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {

	if req.Email == "" || req.Password == "" || req.PasswordConfirm == "" {
		return nil, errors.New("email or password or passwordConfirm is empty")
	}
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("password not match")
	}
	// 生成一个基于密码的哈希值
	// 两个参数：一个是密码，另一个是工作因子（cost）。这里的 cost 参数决定了 bcrypt 哈希算法的计算复杂度, 默认是10
	PasswordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(PasswordHashed),
	}
	err = model.Create(mysql.DB, newUser)
	if err != nil {
		return nil, err
	}
	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
