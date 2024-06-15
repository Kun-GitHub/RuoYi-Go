package input

import (
	"RuoYi-Go/internal/domain/model"
)

// AuthService 输入端口接口
type AuthService interface {
	Login(l model.LoginRequest) (*model.LoginSuccess, error)
	Logout(token string) error
}
