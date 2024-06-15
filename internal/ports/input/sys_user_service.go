package input

import (
	"RuoYi-Go/internal/domain/model"
)

// SysUserService 输入端口接口
type SysUserService interface {
	QueryUserByUserName(username string) (*model.SysUser, error)
}
